// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bladedancer/ratelimit/src/utils (interfaces: TimeSource,JitterRandSource)

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTimeSource is a mock of TimeSource interface
type MockTimeSource struct {
	ctrl     *gomock.Controller
	recorder *MockTimeSourceMockRecorder
}

// MockTimeSourceMockRecorder is the mock recorder for MockTimeSource
type MockTimeSourceMockRecorder struct {
	mock *MockTimeSource
}

// NewMockTimeSource creates a new mock instance
func NewMockTimeSource(ctrl *gomock.Controller) *MockTimeSource {
	mock := &MockTimeSource{ctrl: ctrl}
	mock.recorder = &MockTimeSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimeSource) EXPECT() *MockTimeSourceMockRecorder {
	return m.recorder
}

// UnixNow mocks base method
func (m *MockTimeSource) UnixNow() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnixNow")
	ret0, _ := ret[0].(int64)
	return ret0
}

// UnixNow indicates an expected call of UnixNow
func (mr *MockTimeSourceMockRecorder) UnixNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnixNow", reflect.TypeOf((*MockTimeSource)(nil).UnixNow))
}

// MockJitterRandSource is a mock of JitterRandSource interface
type MockJitterRandSource struct {
	ctrl     *gomock.Controller
	recorder *MockJitterRandSourceMockRecorder
}

// MockJitterRandSourceMockRecorder is the mock recorder for MockJitterRandSource
type MockJitterRandSourceMockRecorder struct {
	mock *MockJitterRandSource
}

// NewMockJitterRandSource creates a new mock instance
func NewMockJitterRandSource(ctrl *gomock.Controller) *MockJitterRandSource {
	mock := &MockJitterRandSource{ctrl: ctrl}
	mock.recorder = &MockJitterRandSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJitterRandSource) EXPECT() *MockJitterRandSourceMockRecorder {
	return m.recorder
}

// Int63 mocks base method
func (m *MockJitterRandSource) Int63() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int63")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Int63 indicates an expected call of Int63
func (mr *MockJitterRandSourceMockRecorder) Int63() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int63", reflect.TypeOf((*MockJitterRandSource)(nil).Int63))
}

// Seed mocks base method
func (m *MockJitterRandSource) Seed(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Seed", arg0)
}

// Seed indicates an expected call of Seed
func (mr *MockJitterRandSourceMockRecorder) Seed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seed", reflect.TypeOf((*MockJitterRandSource)(nil).Seed), arg0)
}
