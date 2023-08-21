// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	store "github.com/tendermint/tendermint/libs/store"
)

// Store is an autogenerated mock type for the Store type
type Store[K comparable, V interface{}] struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *Store[K, V]) All() []V {
	ret := _m.Called()

	var r0 []V
	if rf, ok := ret.Get(0).(func() []V); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]V)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: key
func (_m *Store[K, V]) Delete(key K) {
	_m.Called(key)
}

// Get provides a mock function with given fields: key
func (_m *Store[K, V]) Get(key K) (V, bool) {
	ret := _m.Called(key)

	var r0 V
	var r1 bool
	if rf, ok := ret.Get(0).(func(K) (V, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(K) V); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(V)
	}

	if rf, ok := ret.Get(1).(func(K) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetAndDelete provides a mock function with given fields: key
func (_m *Store[K, V]) GetAndDelete(key K) (V, bool) {
	ret := _m.Called(key)

	var r0 V
	var r1 bool
	if rf, ok := ret.Get(0).(func(K) (V, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(K) V); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(V)
	}

	if rf, ok := ret.Get(1).(func(K) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Put provides a mock function with given fields: key, data
func (_m *Store[K, V]) Put(key K, data V) {
	_m.Called(key, data)
}

// Query provides a mock function with given fields: spec, limit
func (_m *Store[K, V]) Query(spec store.QueryFunc[K, V], limit int) []*V {
	ret := _m.Called(spec, limit)

	var r0 []*V
	if rf, ok := ret.Get(0).(func(store.QueryFunc[K, V], int) []*V); ok {
		r0 = rf(spec, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*V)
		}
	}

	return r0
}

// Update provides a mock function with given fields: key, updates
func (_m *Store[K, V]) Update(key K, updates ...store.UpdateFunc[K, V]) {
	_va := make([]interface{}, len(updates))
	for _i := range updates {
		_va[_i] = updates[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore[K comparable, V interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Store[K, V] {
	mock := &Store[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
