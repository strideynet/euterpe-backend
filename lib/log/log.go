package log

import (
	"context"
	"euterpe/lib/reqid"

	"go.uber.org/zap"
)

func logger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level.SetLevel(zap.DebugLevel)

	logger, _ := config.Build()

	return logger
}

// Package returns a new parent logger. This should be done once per package for "global" style logging, and once per request.
func Package(name string) *zap.Logger {
	return logger().Named(name)
}

// Request returns a parent logger with an included request id tag if one exists in the context
func Request(ctx context.Context) *zap.Logger {
	l := logger().Named("request")
	// Try to find request ID, and add to logger.
	extracted := reqid.Extract(ctx)
	if extracted != "" {
		l = l.With(zap.String("request_id", extracted))
	}

	return l
}
