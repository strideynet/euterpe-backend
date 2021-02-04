// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	teapotv1pb "euterpe/service.teapot/proto"
	mock "github.com/stretchr/testify/mock"
)

// TeapotServer is an autogenerated mock type for the TeapotServer type
type TeapotServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *TeapotServer) Create(_a0 context.Context, _a1 *teapotv1pb.CreateRequest) (*teapotv1pb.CreateResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *teapotv1pb.CreateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *teapotv1pb.CreateRequest) *teapotv1pb.CreateResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*teapotv1pb.CreateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *teapotv1pb.CreateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: _a0, _a1
func (_m *TeapotServer) GetById(_a0 context.Context, _a1 *teapotv1pb.GetByIdRequest) (*teapotv1pb.GetByIdResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *teapotv1pb.GetByIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, *teapotv1pb.GetByIdRequest) *teapotv1pb.GetByIdResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*teapotv1pb.GetByIdResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *teapotv1pb.GetByIdRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedTeapotServer provides a mock function with given fields:
func (_m *TeapotServer) mustEmbedUnimplementedTeapotServer() {
	_m.Called()
}