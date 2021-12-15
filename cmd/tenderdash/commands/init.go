package commands

import (
	"errors"
	"fmt"

	"github.com/dashevo/dashd-go/btcjson"
	"github.com/spf13/cobra"
	cfg "github.com/tendermint/tendermint/config"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

// InitFilesCmd initialises a fresh Tendermint Core instance.
var InitFilesCmd = &cobra.Command{
	Use:       "init [full|validator|seed|single]",
	Short:     "Initializes a Tenderdash node",
	ValidArgs: []string{"full", "validator", "seed", "single"},
	// We allow for zero args so we can throw a more informative error
	Args: cobra.MaximumNArgs(1),
	RunE: initFiles,
}

func initFilesSingleNode(cmd *cobra.Command, args []string) error {
	return initFilesSingleNodeWithConfig(config)
}

var (
	quorumType             int
	coreChainLockedHeight  uint32
	initChainInitialHeight int64
	appHash                []byte
	proTxHash              []byte
)

func AddInitFlags(cmd *cobra.Command) {
	cmd.Flags().IntVar(&quorumType, "quorumType", 0, "Quorum Type")
	cmd.Flags().Uint32Var(&coreChainLockedHeight, "coreChainLockedHeight", 1, "Initial Core Chain Locked Height")
	cmd.Flags().Int64Var(&initChainInitialHeight, "initialHeight", 0, "Initial Height")
	cmd.Flags().BytesHexVar(&proTxHash, "proTxHash", []byte(nil), "Node pro tx hash")
	cmd.Flags().BytesHexVar(&appHash, "appHash", []byte(nil), "App hash")
}

func initFiles(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("must specify a node type: tendermint init [validator|full|seed|single]")
	}
	config.Mode = args[0]
	return initFilesWithConfig(config)
}

func initializeNodeKey(config *cfg.Config) error {
	nodeKeyFile := config.NodeKeyFile()
	if tmos.FileExists(nodeKeyFile) {
		logger.Info("Found node key", "path", nodeKeyFile)
	} else {
		if _, err := p2p.LoadOrGenNodeKey(nodeKeyFile); err != nil {
			return err
		}
		logger.Info("Generated node key", "path", nodeKeyFile)
	}
	return nil
}

func initFilesWithConfig(config *cfg.Config) error {
	var (
		pv  *privval.FilePV
		err error
	)

	if config.Mode == cfg.ModeValidator {
		// private validator
		privValKeyFile := config.PrivValidator.KeyFile()
		privValStateFile := config.PrivValidator.StateFile()
		if tmos.FileExists(privValKeyFile) {
			pv, err = privval.LoadFilePV(privValKeyFile, privValStateFile)
			if err != nil {
				return err
			}

			logger.Info("Found private validator", "keyFile", privValKeyFile,
				"stateFile", privValStateFile)
		} else {
			pv, err = privval.GenFilePV(privValKeyFile, privValStateFile)
			if err != nil {
				return err
			}
			pv.Save()
			logger.Info("Generated private validator", "keyFile", privValKeyFile,
				"stateFile", privValStateFile)
		}
	}

	nodeKeyFile := config.NodeKeyFile()
	if tmos.FileExists(nodeKeyFile) {
		logger.Info("Found node key", "path", nodeKeyFile)
	} else {
		if _, err := types.LoadOrGenNodeKey(nodeKeyFile); err != nil {
			return err
		}
		logger.Info("Generated node key", "path", nodeKeyFile)
	}

	// genesis file
	genFile := config.GenesisFile()
	if tmos.FileExists(genFile) {
		logger.Info("Found genesis file", "path", genFile)
	} else {

		genDoc := types.GenesisDoc{
			ChainID:         fmt.Sprintf("test-chain-%v", tmrand.Str(6)),
			GenesisTime:     tmtime.Now(),
			ConsensusParams: types.DefaultConsensusParams(),
			QuorumType:                   btcjson.LLMQType(quorumType),
			InitialCoreChainLockedHeight: coreChainLockedHeight,
			InitialHeight:                initChainInitialHeight,
			AppHash:                      appHash,
		}
		if keyType == "secp256k1" {
			genDoc.ConsensusParams.Validator = types.ValidatorParams{
				PubKeyTypes: []string{types.ABCIPubKeyTypeSecp256k1},
			}
		}

		ctx, cancel := context.WithTimeout(context.TODO(), ctxTimeout)
		defer cancel()

		// if this is a validator we add it to genesis
		if pv != nil {
			pubKey, err := pv.GetPubKey(ctx)
			if err != nil {
				return fmt.Errorf("can't get pubkey: %w", err)
			}
			genDoc.Validators = []types.GenesisValidator{{
				Address: pubKey.Address(),
				PubKey:  pubKey,
				Power:   10,
			}}
		}

		if err := genDoc.SaveAs(genFile); err != nil {
			return err
		}
		logger.Info("Generated genesis file", "path", genFile)
	}

	// write config file
	if err := cfg.WriteConfigFile(config.RootDir, config); err != nil {
		return err
	}
	logger.Info("Generated config", "mode", config.Mode)

	return nil
}

func initFilesSingleNodeWithConfig(config *cfg.Config) error {
	// private validator
	privValKeyFile := config.PrivValidatorKeyFile()
	privValStateFile := config.PrivValidatorStateFile()
	var pv *privval.FilePV
	if tmos.FileExists(privValKeyFile) {
		pv = privval.LoadFilePV(privValKeyFile, privValStateFile)
		logger.Info("Found private validator", "keyFile", privValKeyFile,
			"stateFile", privValStateFile)
	} else {
		pv = privval.GenFilePV(privValKeyFile, privValStateFile)
		pv.Save()
		logger.Info("Generated private validator", "keyFile", privValKeyFile,
			"stateFile", privValStateFile)
	}

	// node key
	if err := initializeNodeKey(config); err != nil {
		return err
	}

	// genesis file
	genFile := config.GenesisFile()
	if tmos.FileExists(genFile) {
		logger.Info("Found genesis file", "path", genFile)
	} else {
		quorumHash, err := pv.GetFirstQuorumHash()
		if err != nil {
			return fmt.Errorf("there is no quorum hash: %w", err)
		}
		pubKey, err := pv.GetPubKey(quorumHash)
		if err != nil {
			return fmt.Errorf("can't get pubkey in init files with config: %w", err)
		}

		proTxHash, err := pv.GetProTxHash()
		if err != nil {
			return fmt.Errorf("can't get proTxHash: %w", err)
		}

		logger.Info("Found proTxHash", "proTxHash", proTxHash)

		genDoc := types.GenesisDoc{
			ChainID:                      fmt.Sprintf("test-chain-%v", tmrand.Str(6)),
			GenesisTime:                  tmtime.Now(),
			ConsensusParams:              types.DefaultConsensusParams(),
			ThresholdPublicKey:           pubKey,
			QuorumHash:                   quorumHash,
			InitialCoreChainLockedHeight: 1,
		}

		genDoc.Validators = []types.GenesisValidator{{
			PubKey:    pubKey,
			ProTxHash: proTxHash,
			Power:     types.DefaultDashVotingPower,
		}}

		if err := genDoc.SaveAs(genFile); err != nil {
			return err
		}
		logger.Info("Generated genesis file", "path", genFile)
	}

	return nil
}
