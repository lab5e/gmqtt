// Code generated by MockGen. DO NOT EDIT.
// Source: server/service.go

// Package server is a generated GoMock package.
package server

import (
	gomock "github.com/golang/mock/gomock"
	gmqtt "github.com/lab5e/gmqtt"
	session "github.com/lab5e/gmqtt/persistence/session"
	subscription "github.com/lab5e/gmqtt/persistence/subscription"
	retained "github.com/lab5e/gmqtt/retained"
	reflect "reflect"
)

// MockPublisher is a mock of Publisher interface
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockPublisher) Publish(message *gmqtt.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Publish", message)
}

// Publish indicates an expected call of Publish
func (mr *MockPublisherMockRecorder) Publish(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPublisher)(nil).Publish), message)
}

// MockClientService is a mock of ClientService interface
type MockClientService struct {
	ctrl     *gomock.Controller
	recorder *MockClientServiceMockRecorder
}

// MockClientServiceMockRecorder is the mock recorder for MockClientService
type MockClientServiceMockRecorder struct {
	mock *MockClientService
}

// NewMockClientService creates a new mock instance
func NewMockClientService(ctrl *gomock.Controller) *MockClientService {
	mock := &MockClientService{ctrl: ctrl}
	mock.recorder = &MockClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientService) EXPECT() *MockClientServiceMockRecorder {
	return m.recorder
}

// IterateSession mocks base method
func (m *MockClientService) IterateSession(fn session.IterateFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateSession", fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateSession indicates an expected call of IterateSession
func (mr *MockClientServiceMockRecorder) IterateSession(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateSession", reflect.TypeOf((*MockClientService)(nil).IterateSession), fn)
}

// GetSession mocks base method
func (m *MockClientService) GetSession(clientID string) (*gmqtt.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", clientID)
	ret0, _ := ret[0].(*gmqtt.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession
func (mr *MockClientServiceMockRecorder) GetSession(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockClientService)(nil).GetSession), clientID)
}

// GetClient mocks base method
func (m *MockClientService) GetClient(clientID string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient", clientID)
	ret0, _ := ret[0].(Client)
	return ret0
}

// GetClient indicates an expected call of GetClient
func (mr *MockClientServiceMockRecorder) GetClient(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockClientService)(nil).GetClient), clientID)
}

// IterateClient mocks base method
func (m *MockClientService) IterateClient(fn ClientIterateFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateClient", fn)
}

// IterateClient indicates an expected call of IterateClient
func (mr *MockClientServiceMockRecorder) IterateClient(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateClient", reflect.TypeOf((*MockClientService)(nil).IterateClient), fn)
}

// TerminateSession mocks base method
func (m *MockClientService) TerminateSession(clientID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TerminateSession", clientID)
}

// TerminateSession indicates an expected call of TerminateSession
func (mr *MockClientServiceMockRecorder) TerminateSession(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateSession", reflect.TypeOf((*MockClientService)(nil).TerminateSession), clientID)
}

// MockSubscriptionService is a mock of SubscriptionService interface
type MockSubscriptionService struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionServiceMockRecorder
}

// MockSubscriptionServiceMockRecorder is the mock recorder for MockSubscriptionService
type MockSubscriptionServiceMockRecorder struct {
	mock *MockSubscriptionService
}

// NewMockSubscriptionService creates a new mock instance
func NewMockSubscriptionService(ctrl *gomock.Controller) *MockSubscriptionService {
	mock := &MockSubscriptionService{ctrl: ctrl}
	mock.recorder = &MockSubscriptionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubscriptionService) EXPECT() *MockSubscriptionServiceMockRecorder {
	return m.recorder
}

// Subscribe mocks base method
func (m *MockSubscriptionService) Subscribe(clientID string, subscriptions ...*gmqtt.Subscription) (subscription.SubscribeResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{clientID}
	for _, a := range subscriptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(subscription.SubscribeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockSubscriptionServiceMockRecorder) Subscribe(clientID interface{}, subscriptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{clientID}, subscriptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscriptionService)(nil).Subscribe), varargs...)
}

// Unsubscribe mocks base method
func (m *MockSubscriptionService) Unsubscribe(clientID string, topics ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{clientID}
	for _, a := range topics {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unsubscribe", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockSubscriptionServiceMockRecorder) Unsubscribe(clientID interface{}, topics ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{clientID}, topics...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscriptionService)(nil).Unsubscribe), varargs...)
}

// UnsubscribeAll mocks base method
func (m *MockSubscriptionService) UnsubscribeAll(clientID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeAll", clientID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeAll indicates an expected call of UnsubscribeAll
func (mr *MockSubscriptionServiceMockRecorder) UnsubscribeAll(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeAll", reflect.TypeOf((*MockSubscriptionService)(nil).UnsubscribeAll), clientID)
}

// Iterate mocks base method
func (m *MockSubscriptionService) Iterate(fn subscription.IterateFn, options subscription.IterationOptions) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Iterate", fn, options)
}

// Iterate indicates an expected call of Iterate
func (mr *MockSubscriptionServiceMockRecorder) Iterate(fn, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterate", reflect.TypeOf((*MockSubscriptionService)(nil).Iterate), fn, options)
}

// GetStats mocks base method
func (m *MockSubscriptionService) GetStats() subscription.Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats")
	ret0, _ := ret[0].(subscription.Stats)
	return ret0
}

// GetStats indicates an expected call of GetStats
func (mr *MockSubscriptionServiceMockRecorder) GetStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockSubscriptionService)(nil).GetStats))
}

// GetClientStats mocks base method
func (m *MockSubscriptionService) GetClientStats(clientID string) (subscription.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientStats", clientID)
	ret0, _ := ret[0].(subscription.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientStats indicates an expected call of GetClientStats
func (mr *MockSubscriptionServiceMockRecorder) GetClientStats(clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientStats", reflect.TypeOf((*MockSubscriptionService)(nil).GetClientStats), clientID)
}

// MockRetainedService is a mock of RetainedService interface
type MockRetainedService struct {
	ctrl     *gomock.Controller
	recorder *MockRetainedServiceMockRecorder
}

// MockRetainedServiceMockRecorder is the mock recorder for MockRetainedService
type MockRetainedServiceMockRecorder struct {
	mock *MockRetainedService
}

// NewMockRetainedService creates a new mock instance
func NewMockRetainedService(ctrl *gomock.Controller) *MockRetainedService {
	mock := &MockRetainedService{ctrl: ctrl}
	mock.recorder = &MockRetainedServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRetainedService) EXPECT() *MockRetainedServiceMockRecorder {
	return m.recorder
}

// GetRetainedMessage mocks base method
func (m *MockRetainedService) GetRetainedMessage(topicName string) *gmqtt.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRetainedMessage", topicName)
	ret0, _ := ret[0].(*gmqtt.Message)
	return ret0
}

// GetRetainedMessage indicates an expected call of GetRetainedMessage
func (mr *MockRetainedServiceMockRecorder) GetRetainedMessage(topicName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRetainedMessage", reflect.TypeOf((*MockRetainedService)(nil).GetRetainedMessage), topicName)
}

// ClearAll mocks base method
func (m *MockRetainedService) ClearAll() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearAll")
}

// ClearAll indicates an expected call of ClearAll
func (mr *MockRetainedServiceMockRecorder) ClearAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearAll", reflect.TypeOf((*MockRetainedService)(nil).ClearAll))
}

// AddOrReplace mocks base method
func (m *MockRetainedService) AddOrReplace(message *gmqtt.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddOrReplace", message)
}

// AddOrReplace indicates an expected call of AddOrReplace
func (mr *MockRetainedServiceMockRecorder) AddOrReplace(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrReplace", reflect.TypeOf((*MockRetainedService)(nil).AddOrReplace), message)
}

// Remove mocks base method
func (m *MockRetainedService) Remove(topicName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", topicName)
}

// Remove indicates an expected call of Remove
func (mr *MockRetainedServiceMockRecorder) Remove(topicName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockRetainedService)(nil).Remove), topicName)
}

// GetMatchedMessages mocks base method
func (m *MockRetainedService) GetMatchedMessages(topicFilter string) []*gmqtt.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchedMessages", topicFilter)
	ret0, _ := ret[0].([]*gmqtt.Message)
	return ret0
}

// GetMatchedMessages indicates an expected call of GetMatchedMessages
func (mr *MockRetainedServiceMockRecorder) GetMatchedMessages(topicFilter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchedMessages", reflect.TypeOf((*MockRetainedService)(nil).GetMatchedMessages), topicFilter)
}

// Iterate mocks base method
func (m *MockRetainedService) Iterate(fn retained.IterateFn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Iterate", fn)
}

// Iterate indicates an expected call of Iterate
func (mr *MockRetainedServiceMockRecorder) Iterate(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterate", reflect.TypeOf((*MockRetainedService)(nil).Iterate), fn)
}
