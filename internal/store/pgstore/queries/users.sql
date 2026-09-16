-- name: CreateUser :one
INSERT INTO users (user_name, email, password_hash, bio) 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserById :one
SELECT id, user_name, email, password_hash, bio, created_at, updated_at
FROM users 
WHERE id = $1;

-- name: ListUsers :many
SELECT id, user_name, email, bio, created_at, updated_at
FROM users
ORDER BY created_at DESC;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;