package database

import (
	"database/sql"

	"github.com/ricardoraposo/blogs-api-go/internal/entities"
)

type userDB struct {
	*sql.DB
}

func NewUserDB(db *sql.DB) UserDBInterface {
	return &userDB{db}
}

func (a *userDB) GetUsers() ([]entities.User, error) {
	rows, err := a.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []entities.User{}
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.DisplayName, &user.Email, &user.Password, &user.Image); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (a *userDB) CreateUser(user *entities.User) (*entities.User, error) {
	stmt, err := a.Prepare("INSERT INTO users (id, display_name, email, password, image) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(user.ID, user.DisplayName, user.Email, user.Password, user.Image)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *userDB) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	stmt, err := a.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(email).Scan(&user.ID, &user.DisplayName, &user.Email, &user.Password, &user.Image); err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *userDB) GetByID(id string) (*entities.User, error) {
	var user entities.User
	stmt, err := a.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(id).Scan(&user.ID, &user.DisplayName, &user.Email, &user.Password, &user.Image); err != nil {
		return nil, err
	}

	return &user, nil
}
