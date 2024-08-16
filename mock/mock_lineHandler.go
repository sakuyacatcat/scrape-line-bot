// Code generated by MockGen. DO NOT EDIT.
// Source: lineHandler.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	linebot "github.com/line/line-bot-sdk-go/linebot"
)

// MockLineHandler is a mock of LineHandler interface.
type MockLineHandler struct {
	ctrl     *gomock.Controller
	recorder *MockLineHandlerMockRecorder
}

// MockLineHandlerMockRecorder is the mock recorder for MockLineHandler.
type MockLineHandlerMockRecorder struct {
	mock *MockLineHandler
}

// NewMockLineHandler creates a new mock instance.
func NewMockLineHandler(ctrl *gomock.Controller) *MockLineHandler {
	mock := &MockLineHandler{ctrl: ctrl}
	mock.recorder = &MockLineHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLineHandler) EXPECT() *MockLineHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockLineHandler) Handle(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", arg0, arg1)
}

// Handle indicates an expected call of Handle.
func (mr *MockLineHandlerMockRecorder) Handle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockLineHandler)(nil).Handle), arg0, arg1)
}

// MockLineBotClient is a mock of LineBotClient interface.
type MockLineBotClient struct {
	ctrl     *gomock.Controller
	recorder *MockLineBotClientMockRecorder
}

// MockLineBotClientMockRecorder is the mock recorder for MockLineBotClient.
type MockLineBotClientMockRecorder struct {
	mock *MockLineBotClient
}

// NewMockLineBotClient creates a new mock instance.
func NewMockLineBotClient(ctrl *gomock.Controller) *MockLineBotClient {
	mock := &MockLineBotClient{ctrl: ctrl}
	mock.recorder = &MockLineBotClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLineBotClient) EXPECT() *MockLineBotClientMockRecorder {
	return m.recorder
}

// ParseRequest mocks base method.
func (m *MockLineBotClient) ParseRequest(arg0 *http.Request) ([]*linebot.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseRequest", arg0)
	ret0, _ := ret[0].([]*linebot.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseRequest indicates an expected call of ParseRequest.
func (mr *MockLineBotClientMockRecorder) ParseRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseRequest", reflect.TypeOf((*MockLineBotClient)(nil).ParseRequest), arg0)
}

// ReplyMessage mocks base method.
func (m *MockLineBotClient) ReplyMessage(arg0 string, arg1 ...linebot.SendingMessage) *linebot.ReplyMessageCall {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReplyMessage", varargs...)
	ret0, _ := ret[0].(*linebot.ReplyMessageCall)
	return ret0
}

// ReplyMessage indicates an expected call of ReplyMessage.
func (mr *MockLineBotClientMockRecorder) ReplyMessage(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyMessage", reflect.TypeOf((*MockLineBotClient)(nil).ReplyMessage), varargs...)
}

// MockEventHandler is a mock of EventHandler interface.
type MockEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerMockRecorder
}

// MockEventHandlerMockRecorder is the mock recorder for MockEventHandler.
type MockEventHandlerMockRecorder struct {
	mock *MockEventHandler
}

// NewMockEventHandler creates a new mock instance.
func NewMockEventHandler(ctrl *gomock.Controller) *MockEventHandler {
	mock := &MockEventHandler{ctrl: ctrl}
	mock.recorder = &MockEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventHandler) EXPECT() *MockEventHandlerMockRecorder {
	return m.recorder
}

// handleEvent mocks base method.
func (m *MockEventHandler) handleEvent(arg0 *linebot.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleEvent indicates an expected call of handleEvent.
func (mr *MockEventHandlerMockRecorder) handleEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleEvent", reflect.TypeOf((*MockEventHandler)(nil).handleEvent), arg0)
}
