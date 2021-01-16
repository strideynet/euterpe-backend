package domain

import (
	v1 "gmm/service.emoji/proto/v1"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple test of back and fro to ensure nice behaviour
func TestProtoToEmojiAndBack(t *testing.T) {
	proto := &v1.Emoji{
		Id:  "weary",
		Str: "😩",
	}

	v, err := ProtoToEmoji(proto)
	assert.NoError(t, err)
	assert.Equal(t, v, Emoji{
		ID:  proto.Id,
		Str: proto.Str,
	})

	returnProto := v.ToProto()
	assert.Equal(t, proto.String(), returnProto.String())
}
