// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import entity "gitlab.com/konstellation/konstellation-ce/kre/admin-api/domain/entity"
import mock "github.com/stretchr/testify/mock"

// VersionRepo is an autogenerated mock type for the VersionRepo type
type VersionRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: userID, runtimeID, name, description, workflows
func (_m *VersionRepo) Create(userID string, runtimeID string, name string, description string, workflows []entity.Workflow) (*entity.Version, error) {
	ret := _m.Called(userID, runtimeID, name, description, workflows)

	var r0 *entity.Version
	if rf, ok := ret.Get(0).(func(string, string, string, string, []entity.Workflow) *entity.Version); ok {
		r0 = rf(userID, runtimeID, name, description, workflows)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string, []entity.Workflow) error); ok {
		r1 = rf(userID, runtimeID, name, description, workflows)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *VersionRepo) GetByID(id string) (*entity.Version, error) {
	ret := _m.Called(id)

	var r0 *entity.Version
	if rf, ok := ret.Get(0).(func(string) *entity.Version); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRuntime provides a mock function with given fields: runtimeID
func (_m *VersionRepo) GetByRuntime(runtimeID string) ([]entity.Version, error) {
	ret := _m.Called(runtimeID)

	var r0 []entity.Version
	if rf, ok := ret.Get(0).(func(string) []entity.Version); ok {
		r0 = rf(runtimeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(runtimeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: version
func (_m *VersionRepo) Update(version *entity.Version) error {
	ret := _m.Called(version)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Version) error); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
