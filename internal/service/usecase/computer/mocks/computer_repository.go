// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lallison/h_skills_project/internal/service/usecase/computer (interfaces: ComputerRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	slog "log/slog"
	reflect "reflect"

	entities "github.com/lallison/h_skills_project/internal/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockComputerRepository is a mock of ComputerRepository interface.
type MockComputerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockComputerRepositoryMockRecorder
}

// MockComputerRepositoryMockRecorder is the mock recorder for MockComputerRepository.
type MockComputerRepositoryMockRecorder struct {
	mock *MockComputerRepository
}

// NewMockComputerRepository creates a new mock instance.
func NewMockComputerRepository(ctrl *gomock.Controller) *MockComputerRepository {
	mock := &MockComputerRepository{ctrl: ctrl}
	mock.recorder = &MockComputerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputerRepository) EXPECT() *MockComputerRepositoryMockRecorder {
	return m.recorder
}

// ComputerByID mocks base method.
func (m *MockComputerRepository) ComputerByID(arg0 *slog.Logger, arg1 string) (*entities.Computer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputerByID", arg0, arg1)
	ret0, _ := ret[0].(*entities.Computer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputerByID indicates an expected call of ComputerByID.
func (mr *MockComputerRepositoryMockRecorder) ComputerByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputerByID", reflect.TypeOf((*MockComputerRepository)(nil).ComputerByID), arg0, arg1)
}

// Computers mocks base method.
func (m *MockComputerRepository) Computers(arg0 *slog.Logger) ([]*entities.Computer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Computers", arg0)
	ret0, _ := ret[0].([]*entities.Computer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Computers indicates an expected call of Computers.
func (mr *MockComputerRepositoryMockRecorder) Computers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Computers", reflect.TypeOf((*MockComputerRepository)(nil).Computers), arg0)
}

// CreateComputer mocks base method.
func (m *MockComputerRepository) CreateComputer(arg0 *slog.Logger, arg1 *entities.Computer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComputer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateComputer indicates an expected call of CreateComputer.
func (mr *MockComputerRepositoryMockRecorder) CreateComputer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComputer", reflect.TypeOf((*MockComputerRepository)(nil).CreateComputer), arg0, arg1)
}
