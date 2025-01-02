-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    pub_date TIMESTAMP NOT NULL,
    description TEXT NOT NULL,
    link TEXT NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;