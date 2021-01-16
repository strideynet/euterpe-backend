package memory

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gmm/service.emoji/domain"
	"testing"
)

// Nice demonstration of a simple Golang test, avoiding a full blown test table.
// This could be broken down into a test per method, but that seems unnecessary for code this simple.
func TestRepo(t *testing.T) {
	ctx := context.Background()
	t.Parallel()

	t.Run("working empty state", func(t *testing.T) {
		t.Parallel()

		r := New()

		list, err := r.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(list))
	})

	t.Run("add delete findbyid and list functioning", func(t *testing.T) {
		t.Parallel()

		emojis := []domain.Emoji{
			{
				ID:  "corn",
				Str: "🌽",
			},
			{
				ID:  "dog",
				Str: "🦮",
			},
		}

		r := New()

		for _, e := range emojis {
			assert.NoError(t, r.Create(ctx, e))
		}

		list, err := r.List(ctx)
		assert.NoError(t, err)
		assert.Equal(t, len(emojis), len(list))

		// Ensure searching works when exists
		firstEmojiID := emojis[0].ID

		entry, err := r.FindByID(ctx, firstEmojiID)
		assert.NoError(t, err)
		assert.Equal(t, emojis[0].Str, entry.Str)

		err = r.Delete(ctx, firstEmojiID)
		assert.NoError(t, err)

		// Ensure searching works when it does not exist
		entry, err = r.FindByID(ctx, firstEmojiID)
		assert.NoError(t, err)
		assert.Nil(t, entry)
	})
}
