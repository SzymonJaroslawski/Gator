-- name: InsertFeed :one
INSERT INTO feeds (created_at, updated_at, name, URL, user_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetFeedURL :one
SELECT * FROM feeds
WHERE URL = $1;

-- name: MarkFeedFetch :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY  last_fetched_at ASC NULLS FIRST
LIMIT 1;
