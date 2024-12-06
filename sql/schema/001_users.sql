-- +goose Up
CREATE TABLE users (
  id uuid DEFAULT gen_random_uuid(),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  name VARCHAR NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE feeds (
  id uuid DEFAULT gen_random_uuid(),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  name VARCHAR NOT NULL,
  URL VARCHAR UNIQUE NOT NULL,
  user_id uuid NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(user_id) REFERENCES users(id)
  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE users;

DROP TABLE feeds;
