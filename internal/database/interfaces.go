package database

import "github.com/ricardoraposo/blogs-api-go/internal/entities"

type UserDBInterface interface {
	GetUsers() ([]entities.User, error)
    GetByEmail(email string) (*entities.User, error)
    GetByID(id string) (*entities.User, error)
	CreateUser(user *entities.User) (*entities.User, error)
}

type CategoryDBInterface interface {
    CreateCategory(category *entities.Category) (*entities.Category, error)
    GetCategories() ([]*entities.Category, error)
}
