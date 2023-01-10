package iface

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

type IO interface {
	FilterUsers(ctx context.Context) ([]entities.User, error)
	LoginUser(ctx context.Context, user *entities.User) error
	SaveRoom(ctx context.Context, room *entities.Room) error
}
