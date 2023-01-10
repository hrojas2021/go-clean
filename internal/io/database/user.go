package database

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

var (
	selectUsers = `SELECT id,name,username,password,created_at,updated_at FROM users`
)

func (d *Database) FilterUsers(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	err := d.Select(ctx, &users, selectUsers)
	if err != nil {
		return users, err
	}
	return users, nil
}
