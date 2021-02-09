package dao

import (
	"context"
	mockdb "euterpe/lib/db/mock"
	"euterpe/service.teapot/domain"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

var dbMocker *mockdb.Mocker

func TestMain(m *testing.M) {
	dbMocker = mockdb.MustSetup("teapot")
	os.Exit(m.Run())
}

// TODO: Seperate tests here for Create and Find.
func TestDAO_CreateAndFind(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to -short")
	}
	t.Parallel()
	db := dbMocker.DB(t)
	defer db.Close()
	d := DAO{db}
	ctx := context.Background()

	teapot := domain.Teapot{
		ID:         "",
		Name:       "",
		Radius:     0,
		Height:     0,
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}

	err := d.Create(ctx, teapot)
	require.NoError(t, err)

	res, err := d.FindByID(ctx, teapot.ID)
	require.NoError(t, err)
	require.Equal(t, teapot, res)
}
