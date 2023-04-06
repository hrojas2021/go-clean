package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for DB
)

type Database struct {
	DB *sqlx.DB
}

type Args struct {
	Fields  []string
	Join    []string
	Where   []string
	Args    []interface{}
	GroupBy string
}

func New(db *sqlx.DB) *Database {
	return &Database{
		DB: db,
	}
}

func (d *Database) Filter(ctx context.Context, query, defaultFields string, limit, offset uint, total *uint, dest any,
	fns ...func(args *Args)) error {
	args := Args{
		Fields: []string{defaultFields},
	}

	for _, fn := range fns {
		fn(&args)
	}

	var after string
	if len(args.Where) > 0 {
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

	concQuery := fmt.Sprintf(queryWithFields, strings.Join(args.Fields, ", "), after)
	go func() {
		errs <- d.Select(ctx, dest, concQuery, args.Args...)
	}()

	if err := <-errs; err != nil {
		return fmt.Errorf("filter failed; %w; %v", err, <-errs)
	}
	return <-errs
}

func (d *Database) Select(ctx context.Context, dest any, query string, args ...any) error {
	err := d.DB.SelectContext(ctx, dest, query, args...)
	return err
}
