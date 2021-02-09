package dao

import (
	"context"
	"euterpe/service.teapot/domain"
	"github.com/jmoiron/sqlx"
)

type DAO struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *DAO {
	return &DAO{db: db}
}

func (d *DAO) FindByID(ctx context.Context, id string) (*domain.Teapot, error) {
	t := &domain.Teapot{}

	err := d.db.GetContext(ctx, t, "SELECT * FROM teapots WHERE id=$1;", id)
	if err != nil {
		return nil, err
	}

	return t, err
}

func (d *DAO) Create(ctx context.Context, tp domain.Teapot) error {
	t := &domain.Teapot{}

	err := d.db.GetContext(ctx, t, "SELECT * FROM teapots WHERE id=$1;", tp.ID)
	if err != nil {
		return err
	}

	return nil
}
