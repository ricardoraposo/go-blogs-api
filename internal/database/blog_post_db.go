package database

import (
	"database/sql"

	"github.com/ricardoraposo/blogs-api-go/internal/entity"
)

type blogPostDB struct {
	*sql.DB
}

func NewBlogPostDB(db *sql.DB) BlogPostDBInterface {
	return &blogPostDB{db}
}

func (db *blogPostDB) CreateBlogPost(blogPost *entity.BlogPost) (*entity.BlogPost, error) {
	stmt, err := db.Prepare("INSERT INTO blog_posts (id, title, content, user_id, published, updated) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

    _, err = stmt.Exec(blogPost.ID, blogPost.Title, blogPost.Content, blogPost.UserID, blogPost.Published, blogPost.Updated)
    if err != nil {
        return nil, err
    }

    return blogPost, nil
}
