// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	entities "ports_microservice/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// PortStore is an autogenerated mock type for the PortStore type
type PortStore struct {
	mock.Mock
}

type PortStore_Expecter struct {
	mock *mock.Mock
}

func (_m *PortStore) EXPECT() *PortStore_Expecter {
	return &PortStore_Expecter{mock: &_m.Mock}
}

// GetMultiple provides a mock function with given fields:
func (_m *PortStore) GetMultiple() ([]*entities.Port, error) {
	ret := _m.Called()

	var r0 []*entities.Port
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entities.Port, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entities.Port); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Port)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PortStore_GetMultiple_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMultiple'
type PortStore_GetMultiple_Call struct {
	*mock.Call
}

// GetMultiple is a helper method to define mock.On call
func (_e *PortStore_Expecter) GetMultiple() *PortStore_GetMultiple_Call {
	return &PortStore_GetMultiple_Call{Call: _e.mock.On("GetMultiple")}
}

func (_c *PortStore_GetMultiple_Call) Run(run func()) *PortStore_GetMultiple_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PortStore_GetMultiple_Call) Return(_a0 []*entities.Port, _a1 error) *PortStore_GetMultiple_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PortStore_GetMultiple_Call) RunAndReturn(run func() ([]*entities.Port, error)) *PortStore_GetMultiple_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with given fields: port
func (_m *PortStore) Insert(port *entities.Port) error {
	ret := _m.Called(port)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.Port) error); ok {
		r0 = rf(port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PortStore_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type PortStore_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - port *entities.Port
func (_e *PortStore_Expecter) Insert(port interface{}) *PortStore_Insert_Call {
	return &PortStore_Insert_Call{Call: _e.mock.On("Insert", port)}
}

func (_c *PortStore_Insert_Call) Run(run func(port *entities.Port)) *PortStore_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entities.Port))
	})
	return _c
}

func (_c *PortStore_Insert_Call) Return(_a0 error) *PortStore_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PortStore_Insert_Call) RunAndReturn(run func(*entities.Port) error) *PortStore_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// NewPortStore creates a new instance of PortStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPortStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *PortStore {
	mock := &PortStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}