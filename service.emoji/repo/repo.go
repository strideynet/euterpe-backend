package repo

import (
	"context"
	"gmm/service.emoji/domain"
)

// Repo interface allows multiple datasource types to be swapped out
type Repo interface {
	List(ctx context.Context) ([]domain.Emoji, error)

	FindByID(ctx context.Context, id string) (*domain.Emoji, error)

	Create(ctx context.Context, emoji domain.Emoji) error

	Delete(ctx context.Context, id string) error
}
