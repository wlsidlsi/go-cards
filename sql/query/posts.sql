-- name: CreatePost :one
INSERT INTO posts (title, body, user_id, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPost :one
SELECT p.*, u.username as author
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT p.*, u.username as author
FROM posts p
JOIN users u ON p.user_id = u.id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET title = $1, body = $2, status = $3
WHERE id = $4
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;
