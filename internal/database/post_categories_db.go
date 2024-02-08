package database

import "database/sql"

type postCategoryDB struct {
	*sql.DB
}

func NewPostCategoryDB(db *sql.DB) PostCategoryDBInterface {
	return &postCategoryDB{db}
}

func (db *postCategoryDB) CreatePostCategory(postID, categoryID string) error {
    stmt, err := db.Prepare("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)")
    if err != nil {
        return err
    }

    _, err = stmt.Exec(postID, categoryID)
    if err != nil {
        return err
    }

    return nil
}
