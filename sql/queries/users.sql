-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1;

-- name: ResetUser :exec
DELETE FROM users;

-- name: GetAllUsersName :many
SELECT name FROM users;

-- name: GetUserID :one
SELECT * FROM users
WHERE id = $1;
