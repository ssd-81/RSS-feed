-- name: Addfeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds 
WHERE name = $1 LIMIT 1;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;

