-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedSToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST -- null values goes first before any non-null value in ascending order, for can always update the feed that still does not got fetched (NULL value) or the oldest feeds that needs to be updated !
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;
