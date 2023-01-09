package iface

import (
	"context"

	"github.com/google/uuid"
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

type Service interface {
	GetCampaign(ctx context.Context, campaignID uuid.UUID) (*entities.Campaign, error)
}
