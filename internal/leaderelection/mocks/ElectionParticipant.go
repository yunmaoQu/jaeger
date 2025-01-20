// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ElectionParticipant is an autogenerated mock type for the ElectionParticipant type
type ElectionParticipant struct {
	mock.Mock
}

// Close provides a mock function with no fields
func (_m *ElectionParticipant) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsLeader provides a mock function with no fields
func (_m *ElectionParticipant) IsLeader() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsLeader")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Start provides a mock function with no fields
func (_m *ElectionParticipant) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewElectionParticipant creates a new instance of ElectionParticipant. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewElectionParticipant(t interface {
	mock.TestingT
	Cleanup(func())
}) *ElectionParticipant {
	mock := &ElectionParticipant{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}