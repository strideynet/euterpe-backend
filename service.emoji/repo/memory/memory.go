// package memory provides a repository for Emojis backed by a simple slice. This is great for testing where running
// an entire database service is unnecessary.

package memory

import (
	"context"
	errors2 "gmm/lib/errors"
	"gmm/service.emoji/domain"
	"sync"

	"google.golang.org/grpc/codes"
)

type repo struct {
	mu    sync.RWMutex
	items []domain.Emoji
}

// List returns all elements from the repository.
func (r *repo) List(_ context.Context) ([]domain.Emoji, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var res []domain.Emoji

	for i := range r.items {
		res = append(res, r.items[i])
	}

	return res, nil
}

// FindByID returns a single element from the repository, based on its ID, or nil if not found.
func (r *repo) FindByID(_ context.Context, id string) (*domain.Emoji, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Useful for demonstrating an error
	if id == "error" {
		return nil, errors2.Error(codes.DataLoss, "raining cats", map[string]interface{}{
			"cat":    "john",
			"number": 12,
		})
	}

	for i := range r.items {
		if r.items[i].ID == id {
			return &r.items[i], nil
		}
	}

	return nil, nil
}

// Create adds a new element to the repository. TODO: Prevent duplication of id?
func (r *repo) Create(_ context.Context, emoji domain.Emoji) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.items = append(r.items, emoji)

	return nil
}

// Delete removes an element from the repository based on its unique id.
func (r *repo) Delete(_ context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	newItems := make([]domain.Emoji, 0, len(r.items))

	for _, item := range r.items {
		if item.ID != id {
			newItems = append(newItems, item)
		}
	}

	r.items = newItems

	return nil
}

// New instantiates a repo
func New() *repo {
	return &repo{}
}
