package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ricardoraposo/blogs-api-go/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE users (id TEXT, display_name TEXT, email TEXT, password TEXT, image TEXT);")
	if err != nil {
		t.Fatal(err)
	}

	user, err := entity.NewUser("test", "test@test.com", "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	userDB := NewUserDB(db)
	userDB.CreateUser(user)

	var userFromDB entity.User
	stmt, _ := db.Prepare("SELECT * from users where id = ?")
	stmt.QueryRow(user.ID).Scan(&userFromDB.ID, &userFromDB.DisplayName, &userFromDB.Email, &userFromDB.Password, &userFromDB.Image)

	assert.Equal(t, user.ID, userFromDB.ID)
	assert.Equal(t, user.DisplayName, userFromDB.DisplayName)
	assert.Equal(t, user.Password, userFromDB.Password)
	assert.Equal(t, user.Image, userFromDB.Image)
}

func TestGetByEmail(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE users (id TEXT, display_name TEXT, email TEXT, password TEXT, image TEXT);")
	if err != nil {
		t.Fatal(err)
	}

	user, err := entity.NewUser("test", "test@test.com", "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	userDB := NewUserDB(db)
	userDB.CreateUser(user)

    userFromDB, err := userDB.GetByEmail(user.Email)

	assert.Equal(t, user.ID, userFromDB.ID)
	assert.Equal(t, user.DisplayName, userFromDB.DisplayName)
	assert.Equal(t, user.Password, userFromDB.Password)
	assert.Equal(t, user.Image, userFromDB.Image)
}
