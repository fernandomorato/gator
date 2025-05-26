-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name = ?;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;

-- name: TruncateUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;
