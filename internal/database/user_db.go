package database

import (
	"database/sql"

	"github.com/ricardoraposo/blogs-api-go/internal/entity"
)

type userDB struct {
	*sql.DB
}

func NewUserDB(db *sql.DB) UserDBInterface {
	return &userDB{db}
}

func (db *userDB) GetUsers() ([]entity.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []entity.User{}
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.DisplayName, &user.Email, &user.Password, &user.Image); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *userDB) CreateUser(user *entity.User) (*entity.User, error) {
	stmt, err := db.Prepare("INSERT INTO users (id, display_name, email, password, image) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(user.ID, user.DisplayName, user.Email, user.Password, user.Image)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *userDB) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	stmt, err := db.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(email).Scan(&user.ID, &user.DisplayName, &user.Email, &user.Password, &user.Image); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *userDB) GetByID(id string) (*entity.User, error) {
	var user entity.User
	stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(id).Scan(&user.ID, &user.DisplayName, &user.Email, &user.Password, &user.Image); err != nil {
		return nil, err
	}

	return &user, nil
}
