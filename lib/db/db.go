package db

import (
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

// Returns sqlx thats wrapping a pgx connection.
func New(s Settings) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("pgx", fmt.Sprintf("postgres://postgres:devpassword@postgres:5432/%s?sslmode=disable", s.Database))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
