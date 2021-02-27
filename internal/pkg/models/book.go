package models

import (
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type Book struct {
	ID          int64          `db:"id"`
	Title       string         `db:"title"`
	Authors     pq.StringArray `db:"authors"`
	Description sql.NullString `db:"description"`
	Year        sql.NullInt64  `db:"year"`
	Publisher   sql.NullString `db:"publisher"`
	Edition     sql.NullInt64  `db:"edition"`
	ISBN        sql.NullString `db:"isbn"`
	TotalPages  int64          `db:"total_pages"`
	AddedAt     time.Time      `db:"added_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
	UserAddedID int64          `db:"user_added_id"`
}
