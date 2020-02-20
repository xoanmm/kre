// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entity "gitlab.com/konstellation/konstellation-ce/kre/admin-api/domain/entity"
import gql "gitlab.com/konstellation/konstellation-ce/kre/admin-api/adapter/gql"
import mock "github.com/stretchr/testify/mock"

// VersionResolver is an autogenerated mock type for the VersionResolver type
type VersionResolver struct {
	mock.Mock
}

// ConfigurationCompleted provides a mock function with given fields: ctx, obj
func (_m *VersionResolver) ConfigurationCompleted(ctx context.Context, obj *entity.Version) (bool, error) {
	ret := _m.Called(ctx, obj)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Version) bool); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Version) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigurationVariables provides a mock function with given fields: ctx, obj
func (_m *VersionResolver) ConfigurationVariables(ctx context.Context, obj *entity.Version) ([]*gql.ConfigurationVariable, error) {
	ret := _m.Called(ctx, obj)

	var r0 []*gql.ConfigurationVariable
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Version) []*gql.ConfigurationVariable); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gql.ConfigurationVariable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Version) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreationAuthor provides a mock function with given fields: ctx, obj
func (_m *VersionResolver) CreationAuthor(ctx context.Context, obj *entity.Version) (*entity.User, error) {
	ret := _m.Called(ctx, obj)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Version) *entity.User); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Version) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreationDate provides a mock function with given fields: ctx, obj
func (_m *VersionResolver) CreationDate(ctx context.Context, obj *entity.Version) (string, error) {
	ret := _m.Called(ctx, obj)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Version) string); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Version) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicationAuthor provides a mock function with given fields: ctx, obj
func (_m *VersionResolver) PublicationAuthor(ctx context.Context, obj *entity.Version) (*entity.User, error) {
	ret := _m.Called(ctx, obj)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Version) *entity.User); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Version) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PublicationDate provides a mock function with given fields: ctx, obj
func (_m *VersionResolver) PublicationDate(ctx context.Context, obj *entity.Version) (*string, error) {
	ret := _m.Called(ctx, obj)

	var r0 *string
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Version) *string); ok {
		r0 = rf(ctx, obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Version) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: ctx, obj
func (_m *VersionResolver) Status(ctx context.Context, obj *entity.Version) (gql.VersionStatus, error) {
	ret := _m.Called(ctx, obj)

	var r0 gql.VersionStatus
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Version) gql.VersionStatus); ok {
		r0 = rf(ctx, obj)
	} else {
		r0 = ret.Get(0).(gql.VersionStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Version) error); ok {
		r1 = rf(ctx, obj)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}