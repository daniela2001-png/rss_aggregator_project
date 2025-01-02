-- +goose Up
ALTER TABLE posts ADD CONSTRAINT unique_post_link UNIQUE (link);

-- +goose Down
ALTER TABLE posts DROP CONSTRAINT unique_post_link;