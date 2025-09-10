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

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedFromUrl :one
SELECT * FROM feeds 
WHERE url = $1 LIMIT 1;

-- name: GetFeedIdFromUrl :one
SELECT id from feeds 
WHERE url = $1 LIMIT 1;

-- name: GetFeedNameFromFeedId :one
SELECT name from feeds 
WHERE id = $1 LIMIT 1;  

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = $1; 

-- name: GetNextFeedToFetch :one
SELECT * from feeds 
order by last_fetched_at NULLS FIRST LIMIT 1;
