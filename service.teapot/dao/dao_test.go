package dao

import (
	"context"
	mockdb "euterpe/lib/db/mock"
	"euterpe/service.teapot/domain"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

var dbMocker *mockdb.Mocker

func TestMain(m *testing.M) {
	dbMocker = mockdb.New("teapot")
	os.Exit(m.Run())
}

func TestDAO_FindByID(t *testing.T) {
	t.Parallel()
	db := dbMocker.DB(t)
	defer db.Close()
	d := DAO{db}
	ctx := context.Background()
	teapot, err := d.FindByID(ctx, "wew")
	log.Printf("wew %v", teapot)
	assert.NoError(t, err)
	assert.NotNil(t, teapot)
}

func TestDAO_Create(t *testing.T) {
	t.Parallel()
	db := dbMocker.DB(t)
	defer db.Close()
	d := DAO{db}
	ctx := context.Background()

	err := d.Create(ctx, domain.Teapot{
		ID:         "",
		Name:       "",
		Radius:     0,
		Height:     0,
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	})
	assert.NoError(t, err)
}
