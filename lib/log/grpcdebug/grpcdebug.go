// grpcdebug provides a simple Interceptor for emitting the contents of a request and the time taken to process it
package grpcdebug

import (
	"context"
	"gmm/lib/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

// UnaryInterceptor wraps a unary GRPC handler to capture errors and panics
func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	l := log.Request(ctx)
	l.Debug("unary grpc received", zap.String("method", info.FullMethod), zap.Any("req", req))

	start := time.Now()
	defer func() {
		l.Debug("unary grpc handled", zap.Int64("ms", time.Since(start).Milliseconds()))
	}()
	return handler(ctx, req)
}
