package database

import "github.com/ricardoraposo/blogs-api-go/internal/entity"

type UserDBInterface interface {
	GetUsers() ([]entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetByID(id string) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
}

type CategoryDBInterface interface {
	CreateCategory(category *entity.Category) (*entity.Category, error)
	GetCategories() ([]*entity.Category, error)
	HasCategory(id string) (bool, error)
}

type BlogPostDBInterface interface {
	CreateBlogPost(blogPost *entity.BlogPost) (*entity.BlogPost, error)
	GetBlogPosts() ([]*entity.BlogPostWithUser, error)
	GetBlogPostByID(id string) (*entity.BlogPostWithUser, error)
}

type PostCategoryDBInterface interface {
	CreatePostCategory(postID, categoryID string) error
}
