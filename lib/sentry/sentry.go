package sentry

import (
	"context"
	"euterpe/lib/reqid"
	"log"

	"github.com/getsentry/sentry-go"
	"google.golang.org/grpc"
)

type key string

var contextKey key = "sentry"

// TODO: Is init() the cleanest way of doing this? Perhaps sentry should be explicitly setup?
func init() {
	err := sentry.Init(sentry.ClientOptions{})
	if err != nil {
		log.Panicf("sentry.Init: %s", err)
	}
}

func set(ctx context.Context, hub *sentry.Hub) context.Context {
	return context.WithValue(ctx, contextKey, hub)
}

func extract(ctx context.Context) *sentry.Hub {
	val := ctx.Value(contextKey)

	if cast, ok := val.(*sentry.Hub); ok {
		return cast
	}

	return nil
}

// Report extracts a sentry hub from the context, and reports the error
func Report(ctx context.Context, err error) {
	hub := extract(ctx)
	if hub == nil {
		// TODO: log?
		return
	}

	hub = hub.Clone()

	if cast, ok := err.(interface{ Details() map[string]interface{} }); ok {
		details := cast.Details()

		if len(details) > 0 {
			hub.ConfigureScope(func(scope *sentry.Scope) {
				for name, val := range details {
					scope.SetExtra(name, val)
				}
			})
		}

	}

}

// UnaryInterceptor wraps a unary GRPC handler to capture errors and panics
func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	localHub := sentry.CurrentHub().Clone()

	localHub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetExtra("req", req)

		id := reqid.Extract(ctx)
		if id != "" {
			scope.SetExtra("request_id", id)
		}
	})

	ctx = set(ctx, localHub)

	return func() (resp interface{}, err error) {
		defer func() {
			err := recover()
			if err != nil {
				localHub.RecoverWithContext(ctx, err)
			}
		}()

		res, err := handler(ctx, req)
		if err != nil {
			Report(ctx, err)
		}

		return res, err
	}()
}

// StreamInterceptor wraps a stream gRPC handler to capture errors and panics
func StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// TODO: Track errors
	return handler(srv, ss)
}
