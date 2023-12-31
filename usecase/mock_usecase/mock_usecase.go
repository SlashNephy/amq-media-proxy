// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SlashNephy/amq-media-proxy/usecase/media (interfaces: Usecase)
//
// Generated by this command:
//
//	mockgen -typed -package mock_usecase -destination ./mock_usecase/mock_usecase.go github.com/SlashNephy/amq-media-proxy/usecase/media Usecase
//
// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// DownloadMedia mocks base method.
func (m *MockUsecase) DownloadMedia(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadMedia", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadMedia indicates an expected call of DownloadMedia.
func (mr *MockUsecaseMockRecorder) DownloadMedia(arg0, arg1 any) *UsecaseDownloadMediaCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadMedia", reflect.TypeOf((*MockUsecase)(nil).DownloadMedia), arg0, arg1)
	return &UsecaseDownloadMediaCall{Call: call}
}

// UsecaseDownloadMediaCall wrap *gomock.Call
type UsecaseDownloadMediaCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UsecaseDownloadMediaCall) Return(arg0 error) *UsecaseDownloadMediaCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UsecaseDownloadMediaCall) Do(f func(context.Context, string) error) *UsecaseDownloadMediaCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UsecaseDownloadMediaCall) DoAndReturn(f func(context.Context, string) error) *UsecaseDownloadMediaCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FindCachedMediaPath mocks base method.
func (m *MockUsecase) FindCachedMediaPath(arg0 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCachedMediaPath", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FindCachedMediaPath indicates an expected call of FindCachedMediaPath.
func (mr *MockUsecaseMockRecorder) FindCachedMediaPath(arg0 any) *UsecaseFindCachedMediaPathCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCachedMediaPath", reflect.TypeOf((*MockUsecase)(nil).FindCachedMediaPath), arg0)
	return &UsecaseFindCachedMediaPathCall{Call: call}
}

// UsecaseFindCachedMediaPathCall wrap *gomock.Call
type UsecaseFindCachedMediaPathCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UsecaseFindCachedMediaPathCall) Return(arg0 string, arg1 bool) *UsecaseFindCachedMediaPathCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UsecaseFindCachedMediaPathCall) Do(f func(string) (string, bool)) *UsecaseFindCachedMediaPathCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UsecaseFindCachedMediaPathCall) DoAndReturn(f func(string) (string, bool)) *UsecaseFindCachedMediaPathCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
