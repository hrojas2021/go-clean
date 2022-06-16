package campaign

import (
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
	"github.com/hugo.rojas/custom-api/internal/errors"
)

type CampaignRepository interface {
	Noop(ca entities.Campaign) error
	GetCampaignByID(campaign entities.GetCampaignRequest) (models.Campaign, *errors.ErrorResponse)
}
