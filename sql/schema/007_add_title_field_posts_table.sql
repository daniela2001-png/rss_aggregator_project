-- +goose Up
ALTER TABLE posts ADD COLUMN title TEXT NOT NULL;

-- +goose Down
ALTER TABLE posts DROP COLUMN title;