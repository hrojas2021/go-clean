package campaign

import (
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
	"github.com/hugo.rojas/custom-api/internal/errors"
	"github.com/hugo.rojas/custom-api/internal/repository/campaign"
)

type campaignService struct {
	repository campaign.CampaignRepository
}

func NewCampaignService(campaignRepository *campaign.CampaignRepository) CampaignService {
	return &campaignService{
		repository: *campaignRepository,
	}
}

func (s *campaignService) Noop(request models.Campaign) (res models.Campaign) {
	return models.Campaign{}
}

func (s *campaignService) GetCampaign(request models.GetCampaignRequest) (models.Campaign, *errors.ErrorResponse) {
	campaignRequest := entities.GetCampaignRequest{
		ID: request.ID,
	}
	campaign, err := s.repository.GetCampaignByID(campaignRequest)
	// fmt.Printf("TODO BIEN %+v", campaign)
	return campaign, err
}
