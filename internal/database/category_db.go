package database

import (
	"database/sql"

	"github.com/ricardoraposo/blogs-api-go/internal/entities"
)

type categoryDB struct {
	*sql.DB
}

func NewCategoryDB(db *sql.DB) CategoryDBInterface {
	return &categoryDB{db}
}

func (db *categoryDB) CreateCategory(category *entities.Category) (*entities.Category, error) {
	stmt, err := db.Prepare("INSERT INTO categories (id, name) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	if _, err := stmt.Exec(category.ID, category.Name); err != nil {
		return nil, err
	}

	return category, nil
}

func (db *categoryDB) GetCategories() ([]*entities.Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entities.Category
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}
