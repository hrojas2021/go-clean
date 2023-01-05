package database

import "github.com/jmoiron/sqlx"

type Database struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Database {
	return &Database{
		DB: db,
	}
}

func (db *Database) Noop() { // CHECK context
	//Implement functions for this DB
}
