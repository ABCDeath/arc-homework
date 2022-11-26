// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Queue is an autogenerated mock type for the Queue type
type Queue[T interface{}] struct {
	mock.Mock
}

// Dequeue provides a mock function with given fields:
func (_m *Queue[T]) Dequeue() (T, error) {
	ret := _m.Called()

	var r0 T
	if rf, ok := ret.Get(0).(func() T); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(T)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DequeueOrWait provides a mock function with given fields: ctx
func (_m *Queue[T]) DequeueOrWait(ctx context.Context) (T, error) {
	ret := _m.Called(ctx)

	var r0 T
	if rf, ok := ret.Get(0).(func(context.Context) T); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(T)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Enqueue provides a mock function with given fields: object
func (_m *Queue[T]) Enqueue(object T) {
	_m.Called(object)
}

type mockConstructorTestingTNewQueue interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueue creates a new instance of Queue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueue[T interface{}](t mockConstructorTestingTNewQueue) *Queue[T] {
	mock := &Queue[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
