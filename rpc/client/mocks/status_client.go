// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	coretypes "github.com/dashpay/tenderdash/rpc/coretypes"
	mock "github.com/stretchr/testify/mock"
)

// StatusClient is an autogenerated mock type for the StatusClient type
type StatusClient struct {
	mock.Mock
}

// Status provides a mock function with given fields: _a0
func (_m *StatusClient) Status(_a0 context.Context) (*coretypes.ResultStatus, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 *coretypes.ResultStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*coretypes.ResultStatus, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *coretypes.ResultStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretypes.ResultStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStatusClient creates a new instance of StatusClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStatusClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *StatusClient {
	mock := &StatusClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
