// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES (
  $1,
  $2,
  $3
)
RETURNING id, created_at, updated_at, name
`

type CreateUserParams struct {
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Name      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.CreatedAt, arg.UpdatedAt, arg.Name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const getAllUsersName = `-- name: GetAllUsersName :many
SELECT name FROM users
`

func (q *Queries) GetAllUsersName(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsersName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, name FROM users
WHERE name = $1
`

func (q *Queries) GetUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const getUserID = `-- name: GetUserID :one
SELECT id, created_at, updated_at, name FROM users
WHERE id = $1
`

func (q *Queries) GetUserID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const resetUser = `-- name: ResetUser :exec
DELETE FROM users
`

func (q *Queries) ResetUser(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, resetUser)
	return err
}
