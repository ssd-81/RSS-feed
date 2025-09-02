-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
WHERE name = $1 LIMIT 1;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserIDFromName :one
SELECT id from users 
WHERE name = $1 LIMIT 1;

-- name: GetNameFromUserID :one
SELECT name FROM users
WHERE id = $1 LIMIT 1;