//go:generate go run github.com/golang/mock/mockgen -package=mock -source=$GOFILE -destination=mock/service.go
package iface

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
)

type Service interface {
	ListUser(ctx context.Context) ([]entities.User, error)
	Login(context.Context, models.User) (*models.JWT, error)

	SaveRoom(ctx context.Context, room models.Room) (entities.Room, error)
}
