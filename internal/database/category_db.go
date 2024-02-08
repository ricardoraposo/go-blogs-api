package database

import (
	"database/sql"

	"github.com/ricardoraposo/blogs-api-go/internal/entity"
)

type categoryDB struct {
	*sql.DB
}

func NewCategoryDB(db *sql.DB) CategoryDBInterface {
	return &categoryDB{db}
}

func (db *categoryDB) CreateCategory(category *entity.Category) (*entity.Category, error) {
	stmt, err := db.Prepare("INSERT INTO categories (id, name) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	if _, err := stmt.Exec(category.ID, category.Name); err != nil {
		return nil, err
	}

	return category, nil
}

func (db *categoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func (db *categoryDB) HasCategory(id string) (bool, error) {
    var exists bool
    err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id = ?)", id).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}
