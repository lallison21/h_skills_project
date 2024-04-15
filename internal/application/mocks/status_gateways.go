// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lallison/h_skills_project/internal/application (interfaces: StatusGateways)

// Package mocks is a generated GoMock package.
package mocks

import (
	slog "log/slog"
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStatusGateways is a mock of StatusGateways interface.
type MockStatusGateways struct {
	ctrl     *gomock.Controller
	recorder *MockStatusGatewaysMockRecorder
}

// MockStatusGatewaysMockRecorder is the mock recorder for MockStatusGateways.
type MockStatusGatewaysMockRecorder struct {
	mock *MockStatusGateways
}

// NewMockStatusGateways creates a new mock instance.
func NewMockStatusGateways(ctrl *gomock.Controller) *MockStatusGateways {
	mock := &MockStatusGateways{ctrl: ctrl}
	mock.recorder = &MockStatusGatewaysMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusGateways) EXPECT() *MockStatusGatewaysMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockStatusGateways) Status(arg0 *slog.Logger) http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0)
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockStatusGatewaysMockRecorder) Status(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockStatusGateways)(nil).Status), arg0)
}
