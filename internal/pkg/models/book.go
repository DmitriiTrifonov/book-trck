package models

import (
	"database/sql"
	"time"
)

type Book struct {
	ID         int64          `db:"id"`
	Title      string         `db:"title"`
	Year       sql.NullInt64  `db:"year"`
	ISBN       sql.NullString `db:"isbn"`
	TotalPages int64          `db:"total_pages"`
	CreatedAt  time.Time      `db:"created_at"`
}
