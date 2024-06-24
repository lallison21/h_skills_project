// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lallison/h_skills_project/internal/service/usecase/computer (interfaces: ComputerQueueProducer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockComputerQueueProducer is a mock of ComputerQueueProducer interface.
type MockComputerQueueProducer struct {
	ctrl     *gomock.Controller
	recorder *MockComputerQueueProducerMockRecorder
}

// MockComputerQueueProducerMockRecorder is the mock recorder for MockComputerQueueProducer.
type MockComputerQueueProducerMockRecorder struct {
	mock *MockComputerQueueProducer
}

// NewMockComputerQueueProducer creates a new mock instance.
func NewMockComputerQueueProducer(ctrl *gomock.Controller) *MockComputerQueueProducer {
	mock := &MockComputerQueueProducer{ctrl: ctrl}
	mock.recorder = &MockComputerQueueProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputerQueueProducer) EXPECT() *MockComputerQueueProducerMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockComputerQueueProducer) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockComputerQueueProducerMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockComputerQueueProducer)(nil).Error))
}

// Success mocks base method.
func (m *MockComputerQueueProducer) Success() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Success")
	ret0, _ := ret[0].(error)
	return ret0
}

// Success indicates an expected call of Success.
func (mr *MockComputerQueueProducerMockRecorder) Success() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Success", reflect.TypeOf((*MockComputerQueueProducer)(nil).Success))
}