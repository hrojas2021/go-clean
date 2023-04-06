package database

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/errors"
)

var (
	selectUsers = `SELECT id,name,username,password,created_at,updated_at FROM users`
	loginUser   = `SELECT %s FROM users %s`
	loginFields = `id,name,username`
	// `SELECT id,name,username FROM users WHERE username = $1 AND password = $2 LIMIT 1`
)

func (d *Database) FilterUsers(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	err := d.Select(ctx, &users, selectUsers)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (d *Database) LoginUser(ctx context.Context, user *entities.User) error {
	var users []entities.User
	// d.Select(ctx, &users, loginUser, user.Username, user.Password)
	err := d.Filter(ctx, loginUser, loginFields, 1, 0, nil, &users, func(args *Args) {
		args.Where = append(args.Where, `username = $1`)
		args.Where = append(args.Where, `password = $2`)
		args.Args = append(args.Args, user.Username, user.Password)
	})
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return errors.ErrUserNotFound
	}
	return nil
}
