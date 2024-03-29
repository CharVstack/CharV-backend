// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/CharVstack/CharV-backend/usecase/models (interfaces: Socket)

// Package mock_models is a generated GoMock package.
package mock_models

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockSocket is a mock of Socket interface.
type MockSocket struct {
	ctrl     *gomock.Controller
	recorder *MockSocketMockRecorder
}

// MockSocketMockRecorder is the mock recorder for MockSocket.
type MockSocketMockRecorder struct {
	mock *MockSocket
}

// NewMockSocket creates a new mock instance.
func NewMockSocket(ctrl *gomock.Controller) *MockSocket {
	mock := &MockSocket{ctrl: ctrl}
	mock.recorder = &MockSocketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSocket) EXPECT() *MockSocketMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockSocket) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockSocketMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockSocket)(nil).Connect))
}

// Create mocks base method.
func (m *MockSocket) Create(arg0 uuid.UUID) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSocketMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSocket)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockSocket) Delete(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSocketMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSocket)(nil).Delete), arg0)
}

// List mocks base method.
func (m *MockSocket) List() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSocketMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSocket)(nil).List))
}

// SearchFor mocks base method.
func (m *MockSocket) SearchFor(arg0 uuid.UUID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFor", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SearchFor indicates an expected call of SearchFor.
func (mr *MockSocketMockRecorder) SearchFor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFor", reflect.TypeOf((*MockSocket)(nil).SearchFor), arg0)
}

// Send mocks base method.
func (m *MockSocket) Send(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockSocketMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSocket)(nil).Send), arg0)
}
