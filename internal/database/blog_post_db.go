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

func (db *blogPostDB) GetBlogPosts() ([]*entity.BlogPostWithUser, error) {
    stmt, err := db.Prepare("SELECT * FROM blog_posts JOIN users ON blog_posts.user_id = users.id")
    if err != nil {
        return nil, err
    }

    rows, err := stmt.Query()
    if err != nil {
        return nil, err
    }

    var posts []*entity.BlogPostWithUser
    for rows.Next() {
        var post entity.BlogPostWithUser
        err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.Published, &post.Updated, &post.User.ID, &post.User.DisplayName, &post.User.Email)
        if err != nil {
            return nil, err
        }
        posts = append(posts, &post)
    }

    return posts, nil
}
