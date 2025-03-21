// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type Follow struct {
	FollowingUserID sql.NullInt32 `json:"following_user_id"`
	FollowedUserID  sql.NullInt32 `json:"followed_user_id"`
	CreatedAt       sql.NullTime  `json:"created_at"`
}

type Post struct {
	ID    int32          `json:"id"`
	Title sql.NullString `json:"title"`
	// Content of the post
	Body      sql.NullString `json:"body"`
	UserID    sql.NullInt32  `json:"user_id"`
	Status    sql.NullString `json:"status"`
	CreatedAt sql.NullTime   `json:"created_at"`
}

type User struct {
	ID        int32          `json:"id"`
	Username  sql.NullString `json:"username"`
	Role      sql.NullString `json:"role"`
	CreatedAt sql.NullTime   `json:"created_at"`
}
