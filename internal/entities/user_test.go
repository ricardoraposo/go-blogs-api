package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456", "https://gopherbankblobs.s3.amazonaws.com/GOPHER_ROCKS.png")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.DisplayName)
	assert.Equal(t, "j@j.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456", "https://gopherbankblobs.s3.amazonaws.com/GOPHER_ROCKS.png")

	assert.Nil(t, err)
	assert.Nil(t, user.ComparePassword("123456"))
	assert.NotNil(t, user.ComparePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
