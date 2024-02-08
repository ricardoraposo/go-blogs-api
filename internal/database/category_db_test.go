package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ricardoraposo/blogs-api-go/internal/entity"
	"github.com/stretchr/testify/assert"
)

func setupDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE categories (id TEXT PRIMARY KEY, name TEXT)")
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestCreateCategory(t *testing.T) {
	db := setupDB(t)

	categoryDB := NewCategoryDB(db)
	category := entity.NewCategory("test")

	if _, err := categoryDB.CreateCategory(category); err != nil {
		t.Fatal(err)
	}

	var categoryFromDB entity.Category
	err := db.QueryRow("SELECT id, name FROM categories WHERE id = ?", category.ID).Scan(&categoryFromDB.ID, &categoryFromDB.Name)

	assert.Nil(t, err)
	assert.NotNil(t, categoryFromDB)

	assert.Equal(t, category.ID, categoryFromDB.ID)
	assert.Equal(t, category.Name, categoryFromDB.Name)
}

func TestHasCategoryTrue(t *testing.T) {
	db := setupDB(t)

	categoryDB := NewCategoryDB(db)
	category := entity.NewCategory("test")

	categoryDB.CreateCategory(category)

	hasCategory, err := categoryDB.HasCategory(category.ID)
    notHasCategory, err := categoryDB.HasCategory("not-exists")

	assert.Nil(t, err)
	assert.NotNil(t, hasCategory)
	assert.True(t, hasCategory)
    assert.False(t, notHasCategory)
}
