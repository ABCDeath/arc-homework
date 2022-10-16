// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Fuel is an autogenerated mock type for the Fuel type
type Fuel struct {
	mock.Mock
}

// GetFuelAmount provides a mock function with given fields:
func (_m *Fuel) GetFuelAmount() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFuelBurnRate provides a mock function with given fields:
func (_m *Fuel) GetFuelBurnRate() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetFuelAmount provides a mock function with given fields: amount
func (_m *Fuel) SetFuelAmount(amount int) error {
	ret := _m.Called(amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFuel interface {
	mock.TestingT
	Cleanup(func())
}

// NewFuel creates a new instance of Fuel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFuel(t mockConstructorTestingTNewFuel) *Fuel {
	mock := &Fuel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}