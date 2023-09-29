// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SlashNephy/amq-cache-server/usecase/media (interfaces: AMQClient)
//
// Generated by this command:
//
//	mockgen -typed -package mock_repo -destination ./mock_repo/mock_repo.go github.com/SlashNephy/amq-cache-server/usecase/media AMQClient
//
// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAMQClient is a mock of AMQClient interface.
type MockAMQClient struct {
	ctrl     *gomock.Controller
	recorder *MockAMQClientMockRecorder
}

// MockAMQClientMockRecorder is the mock recorder for MockAMQClient.
type MockAMQClientMockRecorder struct {
	mock *MockAMQClient
}

// NewMockAMQClient creates a new mock instance.
func NewMockAMQClient(ctrl *gomock.Controller) *MockAMQClient {
	mock := &MockAMQClient{ctrl: ctrl}
	mock.recorder = &MockAMQClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAMQClient) EXPECT() *MockAMQClientMockRecorder {
	return m.recorder
}

// FetchMedia mocks base method.
func (m *MockAMQClient) FetchMedia(arg0 context.Context, arg1 string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMedia", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMedia indicates an expected call of FetchMedia.
func (mr *MockAMQClientMockRecorder) FetchMedia(arg0, arg1 any) *AMQClientFetchMediaCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMedia", reflect.TypeOf((*MockAMQClient)(nil).FetchMedia), arg0, arg1)
	return &AMQClientFetchMediaCall{Call: call}
}

// AMQClientFetchMediaCall wrap *gomock.Call
type AMQClientFetchMediaCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *AMQClientFetchMediaCall) Return(arg0 *http.Response, arg1 error) *AMQClientFetchMediaCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *AMQClientFetchMediaCall) Do(f func(context.Context, string) (*http.Response, error)) *AMQClientFetchMediaCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *AMQClientFetchMediaCall) DoAndReturn(f func(context.Context, string) (*http.Response, error)) *AMQClientFetchMediaCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
