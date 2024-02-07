-- name: GetUser :one
select * from users where id = ?;

-- name: GetUsers :many
select * from users;

-- name: InsertUser :execresult
insert into users (display_name, email, password, image) values (?, ?, ?, ?);
