package service

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
)

func (s *Service) SaveRoom(ctx context.Context, room *models.Room) error {
	rDB := &entities.Room{
		Name: room.Name,
	}

	err := s.io.SaveRoom(ctx, rDB)
	if err != nil {
		return err
	}
	room.ID = rDB.ID
	room.CreatedAt = rDB.CreatedAt
	room.UpdatedAt = rDB.UpdatedAt

	return nil
}
