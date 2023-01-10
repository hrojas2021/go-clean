package service

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
)

func (s *Service) ListUser(ctx context.Context) ([]entities.User, error) {

	return s.io.FilterUsers(ctx)
}

func (s *Service) Login(ctx context.Context, user models.User) error {
	userEntity := &entities.User{
		Username: user.Username,
		Password: user.Password,
	}

	if err := s.io.LoginUser(ctx, userEntity); err != nil {
		return err
	}
	return nil
}
