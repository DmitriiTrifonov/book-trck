package models

import "github.com/lib/pq"

type User struct {
	ID       int64         `db:"id"`
	Username string        `db:"username"`
	Books    pq.Int64Array `db:"books"`
}
