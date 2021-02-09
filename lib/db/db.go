package db

import (
	"context"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Settings struct {
	Database string
}

func New(ctx context.Context) {
	err, _ := sqlx.ConnectContext(ctx, "pgx", "postgres://postgres:devpassword@postgres:5432/postgres?sslmode=disable")
}
