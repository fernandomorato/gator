-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name as feed_name, users.name as user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = ?;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE feed_follows.user_id = ? AND feed_follows.feed_id = (SELECT id FROM feeds WHERE url = ?);
