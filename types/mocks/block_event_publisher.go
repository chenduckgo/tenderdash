// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	types "github.com/tendermint/tendermint/types"
)

// BlockEventPublisher is an autogenerated mock type for the BlockEventPublisher type
type BlockEventPublisher struct {
	mock.Mock
}

// PublishEventNewBlock provides a mock function with given fields: _a0
func (_m *BlockEventPublisher) PublishEventNewBlock(_a0 types.EventDataNewBlock) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.EventDataNewBlock) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishEventNewBlockHeader provides a mock function with given fields: _a0
func (_m *BlockEventPublisher) PublishEventNewBlockHeader(_a0 types.EventDataNewBlockHeader) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.EventDataNewBlockHeader) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishEventNewEvidence provides a mock function with given fields: _a0
func (_m *BlockEventPublisher) PublishEventNewEvidence(_a0 types.EventDataNewEvidence) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.EventDataNewEvidence) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishEventTx provides a mock function with given fields: _a0
func (_m *BlockEventPublisher) PublishEventTx(_a0 types.EventDataTx) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.EventDataTx) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PublishEventValidatorSetUpdates provides a mock function with given fields: _a0
func (_m *BlockEventPublisher) PublishEventValidatorSetUpdates(_a0 types.EventDataValidatorSetUpdate) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.EventDataValidatorSetUpdate) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlockEventPublisher creates a new instance of BlockEventPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockEventPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockEventPublisher {
	mock := &BlockEventPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
