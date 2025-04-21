-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
)
ON CONFLICT (url) DO UPDATE SET
  updated_at = EXCLUDED.updated_at,
  title = EXCLUDED.title,
  url = EXCLUDED.url,
  description = EXCLUDED.description,
  published_at = EXCLUDED.published_at,
  feed_id = EXCLUDED.feed_id
RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM posts p WHERE p.feed_id IN (SELECT feed_id FROM feed_follows WHERE user_id = $1)
ORDER BY p.published_at DESC
LIMIT $2;
