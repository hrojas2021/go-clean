package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

func (s *Service) GetCampaign(ctx context.Context, campaignID uuid.UUID) (*entities.Campaign, error) {
	c := &entities.Campaign{
		ID: campaignID,
	}

	return c, s.io.GetCampaign(ctx, c)
}
