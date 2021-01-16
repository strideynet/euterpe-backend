// Package reqid assigns a roughly unique id to a requests context, and allows us to fetch it. This gives us a value we
// can we use to search logs/traces with
package reqid

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/xid"
	"google.golang.org/grpc"
)

type key string

var contextKey key = "req_id"

// Set returns a new context with the specified request id set
func Set(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextKey, id)
}

// Extract returns a request id from the context, or "" if it does not exist
func Extract(ctx context.Context) string {
	val := ctx.Value(contextKey)

	if cast, ok := val.(string); ok {
		return cast
	}

	return ""
}

// UnaryInterceptor wraps a unary GRPC handler to capture errors and panics
func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	ctx = Set(ctx, xid.New().String())
	return handler(ctx, req)
}

// StreamInterceptor wraps a stream gRPC handler to capture errors and panics
func StreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	ctx := Set(ss.Context(), xid.New().String())

	wrpd := grpc_middleware.WrapServerStream(ss)
	wrpd.WrappedContext = ctx
	return handler(srv, ss)
}
