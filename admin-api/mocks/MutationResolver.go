// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import gql "gitlab.com/konstellation/konstellation-ce/kre/admin-api/adapter/gql"
import mock "github.com/stretchr/testify/mock"

// MutationResolver is an autogenerated mock type for the MutationResolver type
type MutationResolver struct {
	mock.Mock
}

// ActivateVersion provides a mock function with given fields: ctx, input
func (_m *MutationResolver) ActivateVersion(ctx context.Context, input gql.ActivateVersionInput) (*gql.Version, error) {
	ret := _m.Called(ctx, input)

	var r0 *gql.Version
	if rf, ok := ret.Get(0).(func(context.Context, gql.ActivateVersionInput) *gql.Version); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gql.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, gql.ActivateVersionInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRuntime provides a mock function with given fields: ctx, input
func (_m *MutationResolver) CreateRuntime(ctx context.Context, input gql.CreateRuntimeInput) (*gql.CreateRuntimeResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *gql.CreateRuntimeResponse
	if rf, ok := ret.Get(0).(func(context.Context, gql.CreateRuntimeInput) *gql.CreateRuntimeResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gql.CreateRuntimeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, gql.CreateRuntimeInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateVersion provides a mock function with given fields: ctx, input
func (_m *MutationResolver) CreateVersion(ctx context.Context, input gql.CreateVersionInput) (*gql.CreateVersionResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *gql.CreateVersionResponse
	if rf, ok := ret.Get(0).(func(context.Context, gql.CreateVersionInput) *gql.CreateVersionResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gql.CreateVersionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, gql.CreateVersionInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeployVersion provides a mock function with given fields: ctx, input
func (_m *MutationResolver) DeployVersion(ctx context.Context, input gql.DeployVersionInput) (*gql.Version, error) {
	ret := _m.Called(ctx, input)

	var r0 *gql.Version
	if rf, ok := ret.Get(0).(func(context.Context, gql.DeployVersionInput) *gql.Version); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gql.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, gql.DeployVersionInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSettings provides a mock function with given fields: ctx, input
func (_m *MutationResolver) UpdateSettings(ctx context.Context, input gql.SettingsInput) (*gql.UpdateSettingsResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 *gql.UpdateSettingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, gql.SettingsInput) *gql.UpdateSettingsResponse); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gql.UpdateSettingsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, gql.SettingsInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
