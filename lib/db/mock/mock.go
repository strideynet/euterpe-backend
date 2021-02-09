package mockdb

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // imported for driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // imported for driver
	_ "github.com/jackc/pgx/v4/stdlib"                         // imported for driver
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
)

type Mocker struct {
	svcName string
}

func connString(dbName string) string {
	return fmt.Sprintf("postgres://postgres:devpassword@postgres:5432/%s?sslmode=disable", dbName) // Only runs in docker-compose so we can assume values
}

// DB returns a sqlx.DB instance that wraps a txdb. This should be called for each test, to isolate it's interaction
// with the database and roll back its changes on completion
func (m Mocker) DB(t *testing.T) *sqlx.DB {
	cName := fmt.Sprintf("connection_%d", time.Now().UnixNano())
	conn, err := sql.Open("txdb", cName)
	if err != nil {
		t.Fatalf("Error setting up db: %s", err.Error())
	}

	return sqlx.NewDb(conn, "txdb")
}

// Drops and creates a database with the name provided. Useful for resetting state.
func dropAndCreateDB(dbName string) error {
	conn, err := sqlx.Connect("pgx", connString("postgres")) // connect to default db for creation ops
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec("DROP DATABASE IF EXISTS ?", dbName)
	if err != nil {
		return err
	}

	_, err = conn.Exec("CREATE DATABASE ?", dbName)
	if err != nil {
		return err
	}

	return nil
}

// Setup initiates a new mocker, which creates an appropriate test database if it doesn't exist
// It drops the contents of the database before then running the upwards migration. This lets us test the migrations
// get our database to an expected state. *Setup should be called at most once per test package.*
func Setup(svcName string) (*Mocker, error) {
	dbName := svcName + "_test"
	txdb.Register("txdb", "pgx", connString(dbName))

	err := dropAndCreateDB(dbName)
	if err != nil {
		return nil, err
	}

	m, err := migrate.New(fmt.Sprintf("file:///app/service.%s/migrations", svcName), connString(dbName))
	if err != nil {
		return nil, err
	}
	defer m.Close()
	err = m.Up()
	if err != nil {
		return nil, err
	}

	return &Mocker{svcName: svcName}, nil
}

func MustSetup(svcName string) *Mocker {
	m, err := Setup(svcName)
	if err != nil {
		panic(err)
	}

	return m
}
