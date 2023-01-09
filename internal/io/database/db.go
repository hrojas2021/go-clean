package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Database {
	return &Database{
		DB: db,
	}
}
