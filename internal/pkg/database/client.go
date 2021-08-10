package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DataStorage struct {
	DB *sqlx.DB
}

func NewDataStorage(dsn string) (*DataStorage, error) {
	if dsn == "" {
		return nil, errors.New("dsn is empty")
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DataStorage{
		DB: db,
	}, nil
}
