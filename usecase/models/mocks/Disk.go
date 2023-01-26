// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/CharVstack/CharV-backend/usecase/models (interfaces: Disk)

// Package mock_models is a generated GoMock package.
package mock_models

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockDisk is a mock of Disk interface.
type MockDisk struct {
	ctrl     *gomock.Controller
	recorder *MockDiskMockRecorder
}

// MockDiskMockRecorder is the mock recorder for MockDisk.
type MockDiskMockRecorder struct {
	mock *MockDisk
}

// NewMockDisk creates a new mock instance.
func NewMockDisk(ctrl *gomock.Controller) *MockDisk {
	mock := &MockDisk{ctrl: ctrl}
	mock.recorder = &MockDiskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDisk) EXPECT() *MockDiskMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDisk) Create(arg0 string, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDiskMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDisk)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockDisk) Delete(arg0 string, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDiskMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDisk)(nil).Delete), arg0, arg1)
}
