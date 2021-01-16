package domain

import v1 "gmm/service.emoji/proto/v1"

// Emoji represents a single instance of the Emoji entity
type Emoji struct {
	ID  string
	Str string
}

// ToProto converts emoji to v1 proto Emoji
func (e Emoji) ToProto() *v1.Emoji {
	return &v1.Emoji{
		Id:  e.ID,
		Str: e.Str,
	}
}

// ProtoToEmoji converts a v1 proto Emoji to a domain Emoji
func ProtoToEmoji(emoji *v1.Emoji) (Emoji, error) {
	return Emoji{
		ID:  emoji.GetId(),
		Str: emoji.GetStr(),
	}, nil
}
