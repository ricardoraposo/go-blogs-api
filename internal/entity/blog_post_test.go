package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlogPost(t *testing.T) {
	post := NewBlogPost("Hello, World!", "This is a post", "user1")

	assert.Equal(t, "Hello, World!", post.Title)
	assert.Equal(t, "This is a post", post.Content)
	assert.Equal(t, "user1", post.UserID)
}

func TestNewBlogInvalidTitle(t *testing.T) {
	post := NewBlogPost("", "This is a post", "user1")

	assert.NotNil(t, post.Validate())
	assert.Equal(t, "title is required", post.Validate().Error())
}

func TestNewBlogInvalidContent(t *testing.T) {
	post := NewBlogPost("Hello, World!", "", "user1")

	assert.NotNil(t, post.Validate())
	assert.Equal(t, "content is required", post.Validate().Error())
}

func TestNewBlogInvalidUserID(t *testing.T) {
	post := NewBlogPost("Hello, World!", "This is a post", "")

	assert.NotNil(t, post.Validate())
	assert.Equal(t, "user_id is required", post.Validate().Error())
}

func TestNewBlogTitleTooShort(t *testing.T) {
	post := NewBlogPost("Hi", "This is a post", "user1")

	assert.NotNil(t, post.Validate())
	assert.Equal(t, "title must be at least 5 characters long", post.Validate().Error())
}
