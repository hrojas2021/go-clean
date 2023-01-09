package iface

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

type IO interface {
	GetCampaign(ctx context.Context, campaign *entities.Campaign) error
}
