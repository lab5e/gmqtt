// Code generated by MockGen. DO NOT EDIT.
// Source: persistence/queue/elem.go

// Package queue is a generated GoMock package.
package queue

import (
	gomock "github.com/golang/mock/gomock"
	packets "github.com/lab5e/gmqtt/pkg/packets"
	reflect "reflect"
)

// MockMessageWithID is a mock of MessageWithID interface
type MockMessageWithID struct {
	ctrl     *gomock.Controller
	recorder *MockMessageWithIDMockRecorder
}

// MockMessageWithIDMockRecorder is the mock recorder for MockMessageWithID
type MockMessageWithIDMockRecorder struct {
	mock *MockMessageWithID
}

// NewMockMessageWithID creates a new mock instance
func NewMockMessageWithID(ctrl *gomock.Controller) *MockMessageWithID {
	mock := &MockMessageWithID{ctrl: ctrl}
	mock.recorder = &MockMessageWithIDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageWithID) EXPECT() *MockMessageWithIDMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockMessageWithID) ID() packets.PacketID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(packets.PacketID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockMessageWithIDMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockMessageWithID)(nil).ID))
}

// SetID mocks base method
func (m *MockMessageWithID) SetID(id packets.PacketID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetID", id)
}

// SetID indicates an expected call of SetID
func (mr *MockMessageWithIDMockRecorder) SetID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetID", reflect.TypeOf((*MockMessageWithID)(nil).SetID), id)
}
