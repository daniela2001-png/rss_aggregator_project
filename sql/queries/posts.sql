-- name: CreatePost :exec
INSERT INTO posts (id, title, description, link, pub_date, feed_id)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (link) DO UPDATE -- always updates duplicated values 
SET title = EXCLUDED.title,
    description = EXCLUDED.description,
    pub_date = EXCLUDED.pub_date;

-- name: GetPostsByUserID :many
SELECT title, description, link, pub_date
FROM posts
INNER JOIN feed_follows
ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY pub_date DESC
LIMIT $2;
