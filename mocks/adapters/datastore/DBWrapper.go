// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	datastore "ports_microservice/internal/adapters/datastore"

	mock "github.com/stretchr/testify/mock"
)

// DBWrapper is an autogenerated mock type for the DBWrapper type
type DBWrapper struct {
	mock.Mock
}

type DBWrapper_Expecter struct {
	mock *mock.Mock
}

func (_m *DBWrapper) EXPECT() *DBWrapper_Expecter {
	return &DBWrapper_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: value
func (_m *DBWrapper) Create(value interface{}) datastore.DBWrapper {
	ret := _m.Called(value)

	var r0 datastore.DBWrapper
	if rf, ok := ret.Get(0).(func(interface{}) datastore.DBWrapper); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.DBWrapper)
		}
	}

	return r0
}

// DBWrapper_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type DBWrapper_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - value interface{}
func (_e *DBWrapper_Expecter) Create(value interface{}) *DBWrapper_Create_Call {
	return &DBWrapper_Create_Call{Call: _e.mock.On("Create", value)}
}

func (_c *DBWrapper_Create_Call) Run(run func(value interface{})) *DBWrapper_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *DBWrapper_Create_Call) Return(_a0 datastore.DBWrapper) *DBWrapper_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DBWrapper_Create_Call) RunAndReturn(run func(interface{}) datastore.DBWrapper) *DBWrapper_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields:
func (_m *DBWrapper) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DBWrapper_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type DBWrapper_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
func (_e *DBWrapper_Expecter) Error() *DBWrapper_Error_Call {
	return &DBWrapper_Error_Call{Call: _e.mock.On("Error")}
}

func (_c *DBWrapper_Error_Call) Run(run func()) *DBWrapper_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DBWrapper_Error_Call) Return(_a0 error) *DBWrapper_Error_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DBWrapper_Error_Call) RunAndReturn(run func() error) *DBWrapper_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Find provides a mock function with given fields: value
func (_m *DBWrapper) Find(value interface{}) datastore.DBWrapper {
	ret := _m.Called(value)

	var r0 datastore.DBWrapper
	if rf, ok := ret.Get(0).(func(interface{}) datastore.DBWrapper); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.DBWrapper)
		}
	}

	return r0
}

// DBWrapper_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type DBWrapper_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - value interface{}
func (_e *DBWrapper_Expecter) Find(value interface{}) *DBWrapper_Find_Call {
	return &DBWrapper_Find_Call{Call: _e.mock.On("Find", value)}
}

func (_c *DBWrapper_Find_Call) Run(run func(value interface{})) *DBWrapper_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *DBWrapper_Find_Call) Return(_a0 datastore.DBWrapper) *DBWrapper_Find_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DBWrapper_Find_Call) RunAndReturn(run func(interface{}) datastore.DBWrapper) *DBWrapper_Find_Call {
	_c.Call.Return(run)
	return _c
}

// NewDBWrapper creates a new instance of DBWrapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDBWrapper(t interface {
	mock.TestingT
	Cleanup(func())
}) *DBWrapper {
	mock := &DBWrapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
