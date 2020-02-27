// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import versionpb "gitlab.com/konstellation/kre/k8s-manager/proto/versionpb"

// VersionServiceServer is an autogenerated mock type for the VersionServiceServer type
type VersionServiceServer struct {
	mock.Mock
}

// Publish provides a mock function with given fields: _a0, _a1
func (_m *VersionServiceServer) Publish(_a0 context.Context, _a1 *versionpb.Request) (*versionpb.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *versionpb.Response
	if rf, ok := ret.Get(0).(func(context.Context, *versionpb.Request) *versionpb.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*versionpb.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *versionpb.Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: _a0, _a1
func (_m *VersionServiceServer) Start(_a0 context.Context, _a1 *versionpb.Request) (*versionpb.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *versionpb.Response
	if rf, ok := ret.Get(0).(func(context.Context, *versionpb.Request) *versionpb.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*versionpb.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *versionpb.Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields: _a0, _a1
func (_m *VersionServiceServer) Stop(_a0 context.Context, _a1 *versionpb.Request) (*versionpb.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *versionpb.Response
	if rf, ok := ret.Get(0).(func(context.Context, *versionpb.Request) *versionpb.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*versionpb.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *versionpb.Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unpublish provides a mock function with given fields: _a0, _a1
func (_m *VersionServiceServer) Unpublish(_a0 context.Context, _a1 *versionpb.Request) (*versionpb.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *versionpb.Response
	if rf, ok := ret.Get(0).(func(context.Context, *versionpb.Request) *versionpb.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*versionpb.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *versionpb.Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConfig provides a mock function with given fields: _a0, _a1
func (_m *VersionServiceServer) UpdateConfig(_a0 context.Context, _a1 *versionpb.Request) (*versionpb.Response, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *versionpb.Response
	if rf, ok := ret.Get(0).(func(context.Context, *versionpb.Request) *versionpb.Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*versionpb.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *versionpb.Request) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
