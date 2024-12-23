-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4)
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id;
--

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feeds.user_id = users.id
WHERE feed_follows.user_id = $1;
--

-- name: DeleteFollow :exec
WITH feed_select_id AS (
  SELECT id 
  FROM feeds
  WHERE URL = $1
)
DELETE FROM feed_follows
USING feed_select_id
WHERE feed_follows.feed_id = feed_select_id.id
AND feed_follows.user_id = $2;
