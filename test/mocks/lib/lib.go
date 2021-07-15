// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cri-o/cri-o/pkg/config (interfaces: Iface)

// Package libconfigmock is a generated GoMock package.
package libconfigmock

import (
	reflect "reflect"

	storage "github.com/containers/storage"
	config "github.com/cri-o/cri-o/pkg/config"
	gomock "github.com/golang/mock/gomock"
)

// MockIface is a mock of Iface interface.
type MockIface struct {
	ctrl     *gomock.Controller
	recorder *MockIfaceMockRecorder
}

// MockIfaceMockRecorder is the mock recorder for MockIface.
type MockIfaceMockRecorder struct {
	mock *MockIface
}

// NewMockIface creates a new mock instance.
func NewMockIface(ctrl *gomock.Controller) *MockIface {
	mock := &MockIface{ctrl: ctrl}
	mock.recorder = &MockIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIface) EXPECT() *MockIfaceMockRecorder {
	return m.recorder
}

// GetData mocks base method.
func (m *MockIface) GetData() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// GetData indicates an expected call of GetData.
func (mr *MockIfaceMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockIface)(nil).GetData))
}

// GetStore mocks base method.
func (m *MockIface) GetStore() (storage.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStore")
	ret0, _ := ret[0].(storage.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStore indicates an expected call of GetStore.
func (mr *MockIfaceMockRecorder) GetStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockIface)(nil).GetStore))
}