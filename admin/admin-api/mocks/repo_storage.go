// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method
func (m *MockStorage) CreateBucket(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockStorageMockRecorder) CreateBucket(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockStorage)(nil).CreateBucket), name)
}

// CopyDir mocks base method
func (m *MockStorage) CopyDir(dir, bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyDir", dir, bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyDir indicates an expected call of CopyDir
func (mr *MockStorageMockRecorder) CopyDir(dir, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyDir", reflect.TypeOf((*MockStorage)(nil).CopyDir), dir, bucketName)
}
