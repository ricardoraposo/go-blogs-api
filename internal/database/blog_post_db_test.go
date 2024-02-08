package database

import (
	"database/sql"
	"testing"

	"github.com/ricardoraposo/blogs-api-go/internal/entity"
	"github.com/stretchr/testify/assert"
    _ "github.com/mattn/go-sqlite3"
)

func SetupBlogPostDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE blog_posts (id TEXT PRIMARY KEY, title TEXT, content TEXT, user_id TEXT, published DATETIME, updated DATETIME)")
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestNewBlogPostDB(t *testing.T) {
	db := SetupBlogPostDB(t)

	blogPostDB := NewBlogPostDB(db)
	blogPost := entity.NewBlogPost("title", "content", "user1")

	_, err := blogPostDB.CreateBlogPost(blogPost)

	var blogPostFromDB entity.BlogPost
	db.QueryRow("SELECT * FROM blog_posts WHERE id = ?", blogPost.ID).Scan(&blogPostFromDB.ID, &blogPostFromDB.Title, &blogPostFromDB.Content, &blogPostFromDB.UserID, &blogPostFromDB.Published, &blogPostFromDB.Updated)

	assert.NoError(t, err)
	assert.Equal(t, blogPost.Title, blogPostFromDB.Title)
	assert.Equal(t, blogPost.Content, blogPostFromDB.Content)
	assert.Equal(t, blogPost.UserID, blogPostFromDB.UserID)
	assert.Equal(t, blogPost.ID, blogPostFromDB.ID)
}
