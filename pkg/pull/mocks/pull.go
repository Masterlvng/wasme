// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/wasme/pkg/pull (interfaces: ImagePuller)

// Package mock_pull is a generated GoMock package.
package mock_pull

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/solo-io/wasme/pkg/model"
)

// MockImagePuller is a mock of ImagePuller interface
type MockImagePuller struct {
	ctrl     *gomock.Controller
	recorder *MockImagePullerMockRecorder
}

// MockImagePullerMockRecorder is the mock recorder for MockImagePuller
type MockImagePullerMockRecorder struct {
	mock *MockImagePuller
}

// NewMockImagePuller creates a new mock instance
func NewMockImagePuller(ctrl *gomock.Controller) *MockImagePuller {
	mock := &MockImagePuller{ctrl: ctrl}
	mock.recorder = &MockImagePullerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImagePuller) EXPECT() *MockImagePullerMockRecorder {
	return m.recorder
}

// Pull mocks base method
func (m *MockImagePuller) Pull(arg0 context.Context, arg1 string) (model.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", arg0, arg1)
	ret0, _ := ret[0].(model.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pull indicates an expected call of Pull
func (mr *MockImagePullerMockRecorder) Pull(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockImagePuller)(nil).Pull), arg0, arg1)
}
