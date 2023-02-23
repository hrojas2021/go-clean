package service

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
)

func (s *Service) SaveRoom(ctx context.Context, room *models.Room) error {
	rDb := &entities.Room{
		Name: room.Name,
	}

	err := s.io.SaveRoom(ctx, rDb)
	if err != nil {
		return err
	}
	room.ID = rDb.ID
	room.CreatedAt = rDb.CreatedAt
	room.UpdatedAt = rDb.UpdatedAt

	return nil
}
