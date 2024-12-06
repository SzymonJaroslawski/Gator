-- +goose Up 
SELECT * FROM posts;

-- +goose Down
SELECT * FROM posts;
