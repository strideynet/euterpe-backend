// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package spotifyv1pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SpotifyClient is the client API for Spotify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpotifyClient interface {
	RedeemCode(ctx context.Context, in *RedeemCodeRequest, opts ...grpc.CallOption) (*RedeemCodeResponse, error)
	GetAuthURL(ctx context.Context, in *GetAuthURLRequest, opts ...grpc.CallOption) (*GetAuthURLResponse, error)
	GetListening(ctx context.Context, in *GetListeningRequest, opts ...grpc.CallOption) (*GetListeningResponse, error)
}

type spotifyClient struct {
	cc grpc.ClientConnInterface
}

func NewSpotifyClient(cc grpc.ClientConnInterface) SpotifyClient {
	return &spotifyClient{cc}
}

func (c *spotifyClient) RedeemCode(ctx context.Context, in *RedeemCodeRequest, opts ...grpc.CallOption) (*RedeemCodeResponse, error) {
	out := new(RedeemCodeResponse)
	err := c.cc.Invoke(ctx, "/euterpe.spotify.v1.spotify/RedeemCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyClient) GetAuthURL(ctx context.Context, in *GetAuthURLRequest, opts ...grpc.CallOption) (*GetAuthURLResponse, error) {
	out := new(GetAuthURLResponse)
	err := c.cc.Invoke(ctx, "/euterpe.spotify.v1.spotify/GetAuthURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyClient) GetListening(ctx context.Context, in *GetListeningRequest, opts ...grpc.CallOption) (*GetListeningResponse, error) {
	out := new(GetListeningResponse)
	err := c.cc.Invoke(ctx, "/euterpe.spotify.v1.spotify/GetListening", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpotifyServer is the server API for Spotify service.
// All implementations must embed UnimplementedSpotifyServer
// for forward compatibility
type SpotifyServer interface {
	RedeemCode(context.Context, *RedeemCodeRequest) (*RedeemCodeResponse, error)
	GetAuthURL(context.Context, *GetAuthURLRequest) (*GetAuthURLResponse, error)
	GetListening(context.Context, *GetListeningRequest) (*GetListeningResponse, error)
	mustEmbedUnimplementedSpotifyServer()
}

// UnimplementedSpotifyServer must be embedded to have forward compatible implementations.
type UnimplementedSpotifyServer struct {
}

func (UnimplementedSpotifyServer) RedeemCode(context.Context, *RedeemCodeRequest) (*RedeemCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedeemCode not implemented")
}
func (UnimplementedSpotifyServer) GetAuthURL(context.Context, *GetAuthURLRequest) (*GetAuthURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthURL not implemented")
}
func (UnimplementedSpotifyServer) GetListening(context.Context, *GetListeningRequest) (*GetListeningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListening not implemented")
}
func (UnimplementedSpotifyServer) mustEmbedUnimplementedSpotifyServer() {}

// UnsafeSpotifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpotifyServer will
// result in compilation errors.
type UnsafeSpotifyServer interface {
	mustEmbedUnimplementedSpotifyServer()
}

func RegisterSpotifyServer(s grpc.ServiceRegistrar, srv SpotifyServer) {
	s.RegisterService(&_Spotify_serviceDesc, srv)
}

func _Spotify_RedeemCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedeemCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyServer).RedeemCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/euterpe.spotify.v1.spotify/RedeemCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyServer).RedeemCode(ctx, req.(*RedeemCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Spotify_GetAuthURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyServer).GetAuthURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/euterpe.spotify.v1.spotify/GetAuthURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyServer).GetAuthURL(ctx, req.(*GetAuthURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Spotify_GetListening_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListeningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyServer).GetListening(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/euterpe.spotify.v1.spotify/GetListening",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyServer).GetListening(ctx, req.(*GetListeningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Spotify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "euterpe.spotify.v1.spotify",
	HandlerType: (*SpotifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RedeemCode",
			Handler:    _Spotify_RedeemCode_Handler,
		},
		{
			MethodName: "GetAuthURL",
			Handler:    _Spotify_GetAuthURL_Handler,
		},
		{
			MethodName: "GetListening",
			Handler:    _Spotify_GetListening_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.spotify/proto/spotify.proto",
}
