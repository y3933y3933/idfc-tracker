// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
	"time"
)

type History struct {
	ID        interface{}
	Point     int64
	Reason    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Point struct {
	ID        interface{}
	Total     sql.NullInt64
	Goal      int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
