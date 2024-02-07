// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package blogs

import (
	"context"
	"database/sql"
)

const getUser = `-- name: GetUser :one
select id, display_name, email, password, image from users where id = ?
`

func (q *Queries) GetUser(ctx context.Context, id uint64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Image,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
select id, display_name, email, password, image from users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.DisplayName,
			&i.Email,
			&i.Password,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertUser = `-- name: InsertUser :execresult
insert into users (display_name, email, password, image) values (?, ?, ?, ?)
`

type InsertUserParams struct {
	DisplayName string
	Email       string
	Password    string
	Image       sql.NullString
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertUser,
		arg.DisplayName,
		arg.Email,
		arg.Password,
		arg.Image,
	)
}
