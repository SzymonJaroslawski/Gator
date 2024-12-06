// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
	Name          string
	Url           string
	UserID        uuid.UUID
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

type Post struct {
	ID          uuid.UUID
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	Title       string
	Url         string
	Description string
	PublishedAt time.Time
	FeedID      uuid.UUID
}

type User struct {
	ID        uuid.UUID
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Name      string
}
