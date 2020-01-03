// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import entity "gitlab.com/konstellation/konstellation-ce/kre/admin-api/domain/entity"
import mock "github.com/stretchr/testify/mock"

import time "time"

// VerificationCodeRepo is an autogenerated mock type for the VerificationCodeRepo type
type VerificationCodeRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: code
func (_m *VerificationCodeRepo) Delete(code string) error {
	ret := _m.Called(code)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: code
func (_m *VerificationCodeRepo) Get(code string) (*entity.VerificationCode, error) {
	ret := _m.Called(code)

	var r0 *entity.VerificationCode
	if rf, ok := ret.Get(0).(func(string) *entity.VerificationCode); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.VerificationCode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: code, uid, ttl
func (_m *VerificationCodeRepo) Store(code string, uid string, ttl time.Duration) error {
	ret := _m.Called(code, uid, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, time.Duration) error); ok {
		r0 = rf(code, uid, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
