package database

import (
	"context"
	"errors"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

var (
	selectUsers = `SELECT id,name,username,password,created_at,updated_at FROM users`
	loginUser   = `SELECT id,name,username FROM users WHERE username = $1 AND password = $2 LIMIT 1`
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
	err := d.Select(ctx, &users, loginUser, user.Username, user.Password)

	if err != nil {
		return err
	}

	if len(users) == 0 {
		return errors.New("user not found")
	}
	return err
}
