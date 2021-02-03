package reqid

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestExtract(t *testing.T) {
	ctx := context.Background()
	shouldBeEmptyString := Extract(ctx)

	assert.Equal(t, "", shouldBeEmptyString)
}

func TestUnaryInterceptor(t *testing.T) {
	ctx := context.Background()
	handlerCalled := false

	testRes := "valid"
	res, err := UnaryInterceptor(ctx, struct{}{}, &grpc.UnaryServerInfo{}, func(ctx context.Context, req interface{}) (interface{}, error) {
		handlerCalled = true

		str := Extract(ctx)
		assert.True(t, len(str) > 0)
		return testRes, nil
	})
	assert.NoError(t, err)
	assert.Equal(t, testRes, res)
	assert.True(t, handlerCalled)
}
