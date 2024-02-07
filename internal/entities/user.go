package entities

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         string `json:"id"`
	DisplayName string `json:"display_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Image      string `json:"image"`
}

func NewUser(displayName, email, password, image string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:         uuid.New().String(),
		DisplayName: displayName,
		Email:      email,
		Password:   string(hashedPassword),
		Image:      image,
	}

	return user, nil
}

func (u *User) ComparePassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
