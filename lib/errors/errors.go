// Inspired by Monzo Terrors with a GRPC twist.
// This package wraps the functionality of google.golang.org/grpc/status to add stacktraces and the ability to
// add meta information.
package errors

// TODO: Support GRPC "ErrorDetails"

import (
	"fmt"
	"gmm/lib/errors/stack"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Err struct {
	c          codes.Code
	msg        string
	stackTrace stack.Stack
	details    map[string]interface{}

	cause error
}

// GRPCStatus implemented so that GRPC knows what to return as per google.golang.org/grpc/status FromError()
func (e Err) GRPCStatus() *status.Status {
	return status.New(e.c, e.msg)
}

// Error implemented to meet error interface
func (e Err) Error() string {
	return fmt.Sprintf("[%s] %s", e.c.String(), e.msg)
}

// StackTrace copies and returns the stacktrace for the instantiation of the error.
// Implemented so that Sentry (or other error report tools) are able to determine the trace to report
func (e Err) StackTrace() []uintptr {
	out := make([]uintptr, len(e.stackTrace))
	for i := 0; i < len(e.stackTrace); i++ {
		out[i] = e.stackTrace[i].PC
	}

	return out
}

// Cause returns an error that this one wraps, or nil
func (e Err) Cause() error {
	return e.cause
}

// Exposes our custom details map
func (e Err) Details() map[string]interface{} {
	return e.details
}

func errorFactory(c codes.Code, msg string, details map[string]interface{}) Err {
	return Err{
		c:          c,
		msg:        msg,
		stackTrace: stack.BuildStack(3),
		details:    details,
	}
}

func Error(c codes.Code, format string, a ...interface{}) error {
	// pop off last entry if is map[string]interface{}
	var details map[string]interface{}
	if val, ok := a[len(a)-1].(map[string]interface{}); ok {
		details = val
		a = a[:len(a)-1]
	}

	return errorFactory(c, fmt.Sprintf(format, a...), details)
}

func Wrap(err error, details map[string]interface{}) error {
	e := errorFactory(codes.Internal, err.Error(), details)
	e.cause = err

	return e
}

func WrapWithCode(c codes.Code, err error, details map[string]interface{}) error {
	e := errorFactory(c, err.Error(), details)
	e.cause = err

	return e
}
