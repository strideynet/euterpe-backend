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

// New initiates a new mocker, which creates an appropriate test database if it doesn't exist
// It drops the contents of the database before then running the upwards migration. This lets us test the migrations
// get our database to an expected state.
//
// New should ideally be called once per test package.
func New(svcName string) *Mocker {
	connString := "postgres://postgres:devpassword@postgres:5432/test?sslmode=disable"
	txdb.Register("txdb", "pgx", connString)

	sourceFiles := fmt.Sprintf("file:///app/service.%s/migrations", svcName)

	// Drop any existing content
	m, err := migrate.New(sourceFiles, connString)
	if err != nil {
		panic(err)
	}
	defer m.Close()
	err = m.Drop()
	if err != nil {
		panic(err)
	}

	// Duplicated because: https://github.com/golang-migrate/migrate/issues/226
	m, err = migrate.New(sourceFiles, connString)
	if err != nil {
		panic(err)
	}
	defer m.Close()
	err = m.Up()
	if err != nil {
		panic(err)
	}

	return &Mocker{svcName: svcName}
}
