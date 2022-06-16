package campaign

import (
	"github.com/hugo.rojas/custom-api/internal/domain/models"
	"github.com/hugo.rojas/custom-api/internal/errors"
)

type CampaignService interface {
	Noop(request models.Campaign) (response models.Campaign)
	GetCampaign(request models.GetCampaignRequest) (models.Campaign, *errors.ErrorResponse)
}
