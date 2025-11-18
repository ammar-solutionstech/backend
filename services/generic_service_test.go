package services_test

import (
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"backend/services"
)

func setupTestDB(t *testing.T, modelsToMigrate ...interface{}) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(modelsToMigrate...))
	return db
}

func TestGenericServiceCRUD(t *testing.T) {
	type brandEntity struct {
		ID   int    `gorm:"primaryKey;autoIncrement"`
		Name string `gorm:"size:100;not null"`
	}
	db := setupTestDB(t, &brandEntity{})

	svc := services.NewGenericService[brandEntity](db)

	brand := brandEntity{Name: "Acme"}
	require.NoError(t, svc.Create(&brand))
	require.NotZero(t, brand.ID)

	var list []brandEntity
	require.NoError(t, svc.List(&list))
	require.Len(t, list, 1)

	updated, err := svc.Update(brand.ID, map[string]interface{}{"name": "Acme Updated"})
	require.NoError(t, err)
	require.Equal(t, "Acme Updated", updated.Name)

	require.NoError(t, svc.Delete(brand.ID))
	err = svc.Get(brand.ID, &brandEntity{})
	require.Error(t, err)
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestEnsureHistoryWindow(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour)

	require.NoError(t, services.EnsureHistoryWindow(start, end))
	require.Error(t, services.EnsureHistoryWindow(end, start))
}
