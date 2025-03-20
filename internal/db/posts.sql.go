// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package db

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (title, body, user_id, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, title, body, user_id, status, created_at
`

type CreatePostParams struct {
	Title     sql.NullString `json:"title"`
	Body      sql.NullString `json:"body"`
	UserID    sql.NullInt32  `json:"user_id"`
	Status    sql.NullString `json:"status"`
	CreatedAt sql.NullTime   `json:"created_at"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.Title,
		arg.Body,
		arg.UserID,
		arg.Status,
		arg.CreatedAt,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT p.id, p.title, p.body, p.user_id, p.status, p.created_at, u.username as author
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.id = $1 LIMIT 1
`

type GetPostRow struct {
	ID        int32          `json:"id"`
	Title     sql.NullString `json:"title"`
	Body      sql.NullString `json:"body"`
	UserID    sql.NullInt32  `json:"user_id"`
	Status    sql.NullString `json:"status"`
	CreatedAt sql.NullTime   `json:"created_at"`
	Author    sql.NullString `json:"author"`
}

func (q *Queries) GetPost(ctx context.Context, id int32) (GetPostRow, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i GetPostRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.Author,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT p.id, p.title, p.body, p.user_id, p.status, p.created_at, u.username as author
FROM posts p
JOIN users u ON p.user_id = u.id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2
`

type ListPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListPostsRow struct {
	ID        int32          `json:"id"`
	Title     sql.NullString `json:"title"`
	Body      sql.NullString `json:"body"`
	UserID    sql.NullInt32  `json:"user_id"`
	Status    sql.NullString `json:"status"`
	CreatedAt sql.NullTime   `json:"created_at"`
	Author    sql.NullString `json:"author"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]ListPostsRow, error) {
	rows, err := q.db.QueryContext(ctx, listPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPostsRow
	for rows.Next() {
		var i ListPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
			&i.Author,
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

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET title = $1, body = $2, status = $3
WHERE id = $4
RETURNING id, title, body, user_id, status, created_at
`

type UpdatePostParams struct {
	Title  sql.NullString `json:"title"`
	Body   sql.NullString `json:"body"`
	Status sql.NullString `json:"status"`
	ID     int32          `json:"id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.Title,
		arg.Body,
		arg.Status,
		arg.ID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
