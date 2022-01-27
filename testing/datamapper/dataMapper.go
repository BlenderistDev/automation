// Code generated by MockGen. DO NOT EDIT.
// Source: datamapper/dataMapper.go

// Package mock_datamapper is a generated GoMock package.
package mock_datamapper

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMapping is a mock of Mapping interface.
type MockMapping struct {
	ctrl     *gomock.Controller
	recorder *MockMappingMockRecorder
}

// MockMappingMockRecorder is the mock recorder for MockMapping.
type MockMappingMockRecorder struct {
	mock *MockMapping
}

// NewMockMapping creates a new mock instance.
func NewMockMapping(ctrl *gomock.Controller) *MockMapping {
	mock := &MockMapping{ctrl: ctrl}
	mock.recorder = &MockMappingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMapping) EXPECT() *MockMappingMockRecorder {
	return m.recorder
}

// GetValue mocks base method.
func (m *MockMapping) GetValue() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetValue indicates an expected call of GetValue.
func (mr *MockMappingMockRecorder) GetValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockMapping)(nil).GetValue))
}

// IsSimple mocks base method.
func (m *MockMapping) IsSimple() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSimple")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSimple indicates an expected call of IsSimple.
func (mr *MockMappingMockRecorder) IsSimple() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSimple", reflect.TypeOf((*MockMapping)(nil).IsSimple))
}
