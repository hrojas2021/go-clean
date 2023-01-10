package database

import (
	"context"

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

func (d *Database) Select(ctx context.Context, dest any, query string, args ...any) error {
	err := d.DB.SelectContext(ctx, dest, query, args...)
	return err
}
