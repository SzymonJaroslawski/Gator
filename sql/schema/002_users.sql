-- +goose Up
CREATE TABLE feed_follows (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  feed_id uuid NOT NULL,
  user_id uuid NOT NULL, 
  FOREIGN KEY(feed_id) REFERENCES feeds(id)
  ON DELETE CASCADE,
  FOREIGN KEY(user_id) REFERENCES users(id)
  ON DELETE CASCADE,
  UNIQUE (feed_id, user_id)
);

-- +goose Down
DROP TABLE feed_follows;
