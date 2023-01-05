package extrasdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	sql *sqlx.DB
}
type txContextKey struct{}

type Args struct {
	Fields  []string
	Join    []string
	Where   []string
	Args    []interface{}
	GroupBy string
}

func (d *DB) BeginTransaction(ctx context.Context, callback func(ctx context.Context) error) error {

	var err error
	var tx *sqlx.Tx
	shouldHandle := true

	raw := ctx.Value(txContextKey{})
	if raw != nil {
		tx = raw.(*sqlx.Tx)
		shouldHandle = false
	} else {
		tx, err = d.sql.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}

		ctx = context.WithValue(ctx, txContextKey{}, tx)
	}

	if shouldHandle {
		defer func() {
			if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
				// LOG ERROR
				_ = err
			}
		}()
	}

	if err := callback(ctx); err != nil {
		return err
	}

	if shouldHandle {
		return tx.Commit()
	}

	return nil
}

func (d *DB) Filter(ctx context.Context, query, defaultFields string, limit, offset uint, total *uint, dest any,
	fns ...func(args *Args)) error {

	args := Args{
		Fields: []string{defaultFields},
	}

	for _, fn := range fns {
		fn(&args)
	}

	var after string
	if len(args.Join) > 0 {
		after += fmt.Sprintf(" WHERE %s", strings.Join(args.Where, " AND "))
	}

	if len(args.GroupBy) > 0 {
		after += fmt.Sprintf(" GROUP BY %s", args.GroupBy)
	}

	queryWithFields := query
	var minusArgs int
	if limit > 0 {
		queryWithFields += fmt.Sprintf(" LIMIT $%d", len(args.Args)+1)
		args.Args = append(args.Args, limit)
		minusArgs++
	}

	if offset > 0 {
		queryWithFields += fmt.Sprintf(" OFFSET $%d", len(args.Args)+1)
		args.Args = append(args.Args, offset)
		minusArgs++
	}

	errs := make(chan error)
	go func() {
		var err error
		if total != nil {
			args := args.Args[:len(args.Args)-minusArgs]
			rows := []uint{}
			err = d.Select(ctx, &rows, fmt.Sprintf(query, "COUNT(*) total", after), args...)
			if err == nil {
				if len(rows) == 0 {
					*total = 0
				} else {
					*total = rows[0]
				}
			}
		}
		errs <- err
	}()

	go func() {
		errs <- d.Select(ctx, dest, fmt.Sprintf(queryWithFields, strings.Join(args.Fields, ", "), after), args.Args...)
	}()

	if err := <-errs; err != nil {
		return fmt.Errorf("filter failed; %w; %v", err, <-errs)
	}
	return <-errs
}

func (d *DB) Select(ctx context.Context, dest any, query string, args ...any) error {

	if c := ctx.Value(txContextKey{}); c != nil {
		return c.(*sqlx.Tx).SelectContext(ctx, dest, query, args...)
	}

	err := d.sql.SelectContext(ctx, dest, query, args...)
	return err
}
