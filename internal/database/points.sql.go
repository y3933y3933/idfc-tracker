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

const getPointByUserID = `-- name: GetPointByUserID :one
SELECT id, user_id, total, goal 
FROM points
WHERE user_id = ?
`

type GetPointByUserIDRow struct {
	ID     int64
	UserID int64
	Total  int64
	Goal   int64
}

func (q *Queries) GetPointByUserID(ctx context.Context, userID int64) (GetPointByUserIDRow, error) {
	row := q.db.QueryRowContext(ctx, getPointByUserID, userID)
	var i GetPointByUserIDRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Total,
		&i.Goal,
	)
	return i, err
}

const resetUserPoints = `-- name: ResetUserPoints :exec
UPDATE points SET total = 0 WHERE user_id = ?
`

func (q *Queries) ResetUserPoints(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, resetUserPoints, userID)
	return err
}

const updateGoalByUserID = `-- name: UpdateGoalByUserID :exec
UPDATE points
SET goal = ?
WHERE user_id = ?
`

type UpdateGoalByUserIDParams struct {
	Goal   int64
	UserID int64
}

func (q *Queries) UpdateGoalByUserID(ctx context.Context, arg UpdateGoalByUserIDParams) error {
	_, err := q.db.ExecContext(ctx, updateGoalByUserID, arg.Goal, arg.UserID)
	return err
}

const updateTotalByUserID = `-- name: UpdateTotalByUserID :exec
UPDATE points
SET total = ?
WHERE user_id = ?
`

type UpdateTotalByUserIDParams struct {
	Total  int64
	UserID int64
}

func (q *Queries) UpdateTotalByUserID(ctx context.Context, arg UpdateTotalByUserIDParams) error {
	_, err := q.db.ExecContext(ctx, updateTotalByUserID, arg.Total, arg.UserID)
	return err
}
