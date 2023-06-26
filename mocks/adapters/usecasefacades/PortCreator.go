// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "ports_microservice/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// PortCreator is an autogenerated mock type for the PortCreator type
type PortCreator struct {
	mock.Mock
}

type PortCreator_Expecter struct {
	mock *mock.Mock
}

func (_m *PortCreator) EXPECT() *PortCreator_Expecter {
	return &PortCreator_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, port
func (_m *PortCreator) Execute(ctx context.Context, port *entities.Port) (*entities.Port, error) {
	ret := _m.Called(ctx, port)

	var r0 *entities.Port
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Port) (*entities.Port, error)); ok {
		return rf(ctx, port)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Port) *entities.Port); ok {
		r0 = rf(ctx, port)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Port)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.Port) error); ok {
		r1 = rf(ctx, port)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PortCreator_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type PortCreator_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - port *entities.Port
func (_e *PortCreator_Expecter) Execute(ctx interface{}, port interface{}) *PortCreator_Execute_Call {
	return &PortCreator_Execute_Call{Call: _e.mock.On("Execute", ctx, port)}
}

func (_c *PortCreator_Execute_Call) Run(run func(ctx context.Context, port *entities.Port)) *PortCreator_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.Port))
	})
	return _c
}

func (_c *PortCreator_Execute_Call) Return(_a0 *entities.Port, _a1 error) *PortCreator_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PortCreator_Execute_Call) RunAndReturn(run func(context.Context, *entities.Port) (*entities.Port, error)) *PortCreator_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewPortCreator creates a new instance of PortCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPortCreator(t interface {
	mock.TestingT
	Cleanup(func())
}) *PortCreator {
	mock := &PortCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
