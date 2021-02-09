package db

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Settings struct {
	Database string
}

func ServiceSettingsFromEnv(svcName string) Settings {
	return Settings{Database: svcName}
}

func New(s Settings) (*sqlx.DB, error) {
	err, _ := sqlx.Connect("pgx", fmt.Sprintf("postgres://postgres:devpassword@postgres:5432/%s?sslmode=disable", s.Database))
}
