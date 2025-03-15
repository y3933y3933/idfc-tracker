// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
)

type Config struct {
	Key   string
	Value string
}

type History struct {
	ID        int64
	UserID    int64
	Point     int64
	Reason    string
	CreatedAt sql.NullTime
}

type Point struct {
	ID        int64
	UserID    int64
	Total     sql.NullInt64
	Goal      int64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type User struct {
	ID        int64
	Name      string
	CreatedAt sql.NullTime
}
