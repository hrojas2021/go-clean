package service

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

func (s *Service) ListUser(ctx context.Context) ([]entities.User, error) {

	return s.io.FilterUsers(ctx)
}
