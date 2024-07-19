// Code generated by MockGen. DO NOT EDIT.
// Source: internal/config/kod_gen_interface.go
//
// Generated by this command:
//
//	mockgen -source internal/config/kod_gen_interface.go -destination internal/config/kod_gen_mock.go -package config
//

// Package config is a generated GoMock package.
package config

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockConfig is a mock of Config interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig.
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockConfig) Config() *ConfigInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*ConfigInfo)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockConfigMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockConfig)(nil).Config))
}