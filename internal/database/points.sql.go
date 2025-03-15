// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: points.sql

package database

import (
	"context"
)

const createPoint = `-- name: CreatePoint :exec
INSERT INTO points (user_id, goal, created_at)
VALUES (?, ?, CURRENT_TIMESTAMP)
`

type CreatePointParams struct {
	UserID int64
	Goal   int64
}

func (q *Queries) CreatePoint(ctx context.Context, arg CreatePointParams) error {
	_, err := q.db.ExecContext(ctx, createPoint, arg.UserID, arg.Goal)
	return err
}
