package models

import (
	"database/sql"
	"time"
)

type Author struct {
	ID         int64          `db:"id"`
	FirstName  string         `db:"first_name"`
	MiddleName sql.NullString `db:"middle_name"`
	LastName   sql.NullString `db:"last_name"`
	CreatedAt  time.Time      `db:"created_at"`
}
