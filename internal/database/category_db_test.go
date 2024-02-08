package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ricardoraposo/blogs-api-go/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

    db.Exec("CREATE TABLE categories (id TEXT PRIMARY KEY, name TEXT)")

	categoryDB := NewCategoryDB(db)
	category := entities.NewCategory("test")

	if _, err := categoryDB.CreateCategory(category); err != nil {
		t.Fatal(err)
	}

	var categoryFromDB entities.Category
	err = db.QueryRow("SELECT id, name FROM categories WHERE id = ?", category.ID).Scan(&categoryFromDB.ID, &categoryFromDB.Name)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, categoryFromDB)
	assert.Nil(t, err)

    assert.Equal(t, category.ID, categoryFromDB.ID)
    assert.Equal(t, category.Name, categoryFromDB.Name)
}
