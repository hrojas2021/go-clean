package service

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
)

func (s *Service) SaveRoom(ctx context.Context, room models.Room) (entities.Room, error) {
	r := &entities.Room{
		Name: room.Name,
	}

	err := s.io.SaveRoom(ctx, r)
	if err != nil {
		return entities.Room{}, err
	}
	return *r, nil
}
