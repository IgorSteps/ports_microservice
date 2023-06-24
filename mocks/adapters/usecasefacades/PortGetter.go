// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	entities "ports_microservice/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// PortGetter is an autogenerated mock type for the PortGetter type
type PortGetter struct {
	mock.Mock
}

type PortGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *PortGetter) EXPECT() *PortGetter_Expecter {
	return &PortGetter_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields:
func (_m *PortGetter) Execute() ([]*entities.Port, error) {
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

// PortGetter_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type PortGetter_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *PortGetter_Expecter) Execute() *PortGetter_Execute_Call {
	return &PortGetter_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *PortGetter_Execute_Call) Run(run func()) *PortGetter_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PortGetter_Execute_Call) Return(_a0 []*entities.Port, _a1 error) *PortGetter_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PortGetter_Execute_Call) RunAndReturn(run func() ([]*entities.Port, error)) *PortGetter_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewPortGetter creates a new instance of PortGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPortGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *PortGetter {
	mock := &PortGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
