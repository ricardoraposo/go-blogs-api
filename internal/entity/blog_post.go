package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTitleIsRequired = errors.New("title is required")
    ErrTitleTooShort = errors.New("title must be at least 5 characters long")
    ErrContentIsRequired = errors.New("content is required")
    ErrContentTooShort = errors.New("title must be at least 5 characters long")
    ErrUserIDIsRequired = errors.New("user_id is required")
)

type BlogPost struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    string    `json:"user_id"`
	Published time.Time `json:"published"`
	Updated   time.Time `json:"updated"`
}

func NewBlogPost(title string, content string, userID string) *BlogPost {
	return &BlogPost{
		ID:        uuid.New().String(),
		Title:     title,
		Content:   content,
		UserID:    userID,
		Published: time.Now(),
		Updated:   time.Now(),
	}
}

func (b *BlogPost) Validate() error {
	if b.Title == "" {
		return ErrTitleIsRequired
	}
    if len(b.Title) < 5 {
        return ErrTitleTooShort
    }
	if b.Content == "" {
		return ErrContentIsRequired
	}
    if len(b.Content) < 5 {
        return ErrContentTooShort
    }
	if b.UserID == "" {
		return ErrUserIDIsRequired
	}
	return nil
}
