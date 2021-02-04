// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	spotifyv1pb "euterpe/service.spotify/proto"
)

// SpotifyClient is an autogenerated mock type for the SpotifyClient type
type SpotifyClient struct {
	mock.Mock
}

// GetAuthURL provides a mock function with given fields: ctx, in, opts
func (_m *SpotifyClient) GetAuthURL(ctx context.Context, in *spotifyv1pb.GetAuthURLRequest, opts ...grpc.CallOption) (*spotifyv1pb.GetAuthURLResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *spotifyv1pb.GetAuthURLResponse
	if rf, ok := ret.Get(0).(func(context.Context, *spotifyv1pb.GetAuthURLRequest, ...grpc.CallOption) *spotifyv1pb.GetAuthURLResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spotifyv1pb.GetAuthURLResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *spotifyv1pb.GetAuthURLRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListening provides a mock function with given fields: ctx, in, opts
func (_m *SpotifyClient) GetListening(ctx context.Context, in *spotifyv1pb.GetListeningRequest, opts ...grpc.CallOption) (*spotifyv1pb.GetListeningResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *spotifyv1pb.GetListeningResponse
	if rf, ok := ret.Get(0).(func(context.Context, *spotifyv1pb.GetListeningRequest, ...grpc.CallOption) *spotifyv1pb.GetListeningResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spotifyv1pb.GetListeningResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *spotifyv1pb.GetListeningRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RedeemCode provides a mock function with given fields: ctx, in, opts
func (_m *SpotifyClient) RedeemCode(ctx context.Context, in *spotifyv1pb.RedeemCodeRequest, opts ...grpc.CallOption) (*spotifyv1pb.RedeemCodeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *spotifyv1pb.RedeemCodeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *spotifyv1pb.RedeemCodeRequest, ...grpc.CallOption) *spotifyv1pb.RedeemCodeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spotifyv1pb.RedeemCodeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *spotifyv1pb.RedeemCodeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
