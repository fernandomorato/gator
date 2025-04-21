-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
  ) RETURNING *
)

SELECT ff.*, u.name as user_name, f.name as feed_name
FROM inserted_feed_follow ff
INNER JOIN users u ON u.id = ff.user_id
INNER JOIN feeds f ON f.id = ff.feed_id;
