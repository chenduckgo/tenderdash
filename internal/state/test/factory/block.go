package factory

import (
	"context"
	"encoding/binary"
	"fmt"

	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"
	sm "github.com/tendermint/tendermint/internal/state"
	"github.com/tendermint/tendermint/internal/test/factory"
	"github.com/tendermint/tendermint/types"
)

func MakeBlocks(ctx context.Context, t *testing.T, n int, state *sm.State, privVal types.PrivValidator) []*types.Block {
	t.Helper()

	blocks := make([]*types.Block, n)

	var (
		prevBlock     *types.Block
		prevBlockMeta *types.BlockMeta
	)

	appHeight := byte(0x01)
	for i := 0; i < n; i++ {
		height := int64(i + 1)

		block, parts := makeBlockAndPartSet(ctx, t, *state, prevBlock, prevBlockMeta, privVal, height)
		blocks = append(blocks, block)

		prevBlock = block
		prevBlockMeta = types.NewBlockMeta(block, parts)

		// update state
		state.LastStateID = state.StateID()
		state.AppHash = make([]byte, crypto.DefaultAppHashSize)
		binary.BigEndian.PutUint64(state.AppHash, uint64(height))
		appHeight++
		state.LastBlockHeight = height
	}

	return blocks
}

func MakeBlock(state sm.State, height int64, c *types.Commit, coreChainLock *types.CoreChainLock, proposedAppVersion uint64) (*types.Block, error) {
	if state.LastBlockHeight != (height - 1) {
		return nil, fmt.Errorf("requested height %d should be 1 more than last block height %d",
			height, state.LastBlockHeight)
	}
	block, _ :=  state.MakeBlock(
		height,
		coreChainLock,
		factory.MakeNTxs(state.LastBlockHeight, 10),
		c,
		nil,
		state.Validators.GetProposer().ProTxHash,
		proposedAppVersion,
	)
	return block, nil
}

func makeBlockAndPartSet(
	ctx context.Context,
	t *testing.T,
	state sm.State,
	lastBlock *types.Block,
	lastBlockMeta *types.BlockMeta,
	privVal types.PrivValidator,
	height int64,
) (*types.Block, *types.PartSet) {
	t.Helper()

	lastCommit := types.NewCommit(height-1, 0, types.BlockID{}, nil)
	if height > 1 {
		vote, err := factory.MakeVote(
			ctx,
			privVal,
			state.Validators,
			lastBlock.Header.ChainID,
			1, lastBlock.Header.Height, 0, 2,
			lastBlockMeta.BlockID,
			state.LastStateID,
			time.Now())
		require.NoError(t, err)
		lastCommit = types.NewCommit(vote.Height, vote.Round,
			lastBlockMeta.BlockID, state.StateID(), state.LastValidators.QuorumHash, vote.BlockSignature, vote.StateSignature)
	}

	block := state.MakeBlock(height, []types.Tx{}, lastCommit, nil, state.Validators.GetProposer().Address)
	partSet, err := block.MakePartSet(types.BlockPartSizeBytes)
	require.NoError(t, err)

	return block, partSet
}
