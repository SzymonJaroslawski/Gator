-- +goose Up
CREATE TABLE posts (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  title VARCHAR NOT NULL,
  URL VARCHAR NOT NULL,
  description VARCHAR NOT NULL,
  published_at TIMESTAMP NOT NULL,
  feed_id uuid NOT NULL,
  FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
  UNIQUE(URL)
);

-- +goose Down 
DROP TABLE posts;
