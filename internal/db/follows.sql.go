// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: follows.sql

package db

import (
	"context"
	"database/sql"
)

const followUser = `-- name: FollowUser :exec
INSERT INTO follows (following_user_id, followed_user_id, created_at)
VALUES ($1, $2, now())
`

type FollowUserParams struct {
	FollowingUserID sql.NullInt32 `json:"following_user_id"`
	FollowedUserID  sql.NullInt32 `json:"followed_user_id"`
}

func (q *Queries) FollowUser(ctx context.Context, arg FollowUserParams) error {
	_, err := q.db.ExecContext(ctx, followUser, arg.FollowingUserID, arg.FollowedUserID)
	return err
}

const getFollowers = `-- name: GetFollowers :many
SELECT f.following_user_id, f.followed_user_id, f.created_at, u.username as follower_username
FROM follows f
JOIN users u ON f.following_user_id = u.id
WHERE f.followed_user_id = $1
`

type GetFollowersRow struct {
	FollowingUserID  sql.NullInt32  `json:"following_user_id"`
	FollowedUserID   sql.NullInt32  `json:"followed_user_id"`
	CreatedAt        sql.NullTime   `json:"created_at"`
	FollowerUsername sql.NullString `json:"follower_username"`
}

func (q *Queries) GetFollowers(ctx context.Context, followedUserID sql.NullInt32) ([]GetFollowersRow, error) {
	rows, err := q.db.QueryContext(ctx, getFollowers, followedUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFollowersRow
	for rows.Next() {
		var i GetFollowersRow
		if err := rows.Scan(
			&i.FollowingUserID,
			&i.FollowedUserID,
			&i.CreatedAt,
			&i.FollowerUsername,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollows = `-- name: GetFollows :many
SELECT f.following_user_id, f.followed_user_id, f.created_at, u.username as followed_username
FROM follows f
JOIN users u ON f.followed_user_id = u.id
WHERE f.following_user_id = $1
`

type GetFollowsRow struct {
	FollowingUserID  sql.NullInt32  `json:"following_user_id"`
	FollowedUserID   sql.NullInt32  `json:"followed_user_id"`
	CreatedAt        sql.NullTime   `json:"created_at"`
	FollowedUsername sql.NullString `json:"followed_username"`
}

func (q *Queries) GetFollows(ctx context.Context, followingUserID sql.NullInt32) ([]GetFollowsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFollows, followingUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFollowsRow
	for rows.Next() {
		var i GetFollowsRow
		if err := rows.Scan(
			&i.FollowingUserID,
			&i.FollowedUserID,
			&i.CreatedAt,
			&i.FollowedUsername,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unfollowUser = `-- name: UnfollowUser :exec
DELETE FROM follows
WHERE following_user_id = $1 AND followed_user_id = $2
`

type UnfollowUserParams struct {
	FollowingUserID sql.NullInt32 `json:"following_user_id"`
	FollowedUserID  sql.NullInt32 `json:"followed_user_id"`
}

func (q *Queries) UnfollowUser(ctx context.Context, arg UnfollowUserParams) error {
	_, err := q.db.ExecContext(ctx, unfollowUser, arg.FollowingUserID, arg.FollowedUserID)
	return err
}
