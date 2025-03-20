-- name: FollowUser :exec
INSERT INTO follows (following_user_id, followed_user_id, created_at)
VALUES ($1, $2, now());

-- name: UnfollowUser :exec
DELETE FROM follows
WHERE following_user_id = $1 AND followed_user_id = $2;

-- name: GetFollows :many
SELECT f.*, u.username as followed_username
FROM follows f
JOIN users u ON f.followed_user_id = u.id
WHERE f.following_user_id = $1;

-- name: GetFollowers :many
SELECT f.*, u.username as follower_username
FROM follows f
JOIN users u ON f.following_user_id = u.id
WHERE f.followed_user_id = $1;