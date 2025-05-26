-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetFeedByUrl :one
SELECT * FROM feeds where url = ?;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds SET updated_at = NOW(), last_fetched_at = NOW() WHERE id = ?;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;
