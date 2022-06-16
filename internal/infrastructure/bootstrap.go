package infrastructure

import (
	"github.com/hugo.rojas/custom-api/internal/controllers"
	repository "github.com/hugo.rojas/custom-api/internal/repository/campaign"
	service "github.com/hugo.rojas/custom-api/internal/service/campaign"
)

type Bootstraper interface {
	startNoopService()
	startCampaignService()
}

type bootstrap struct {
	api *API
}

func NewBootstrap(a *API) Bootstraper {
	return &bootstrap{
		api: a,
	}
}

func InitServices(b Bootstraper) {
	b.startNoopService()
	b.startCampaignService()
}

func (b *bootstrap) startNoopService() {
	initNoopRoutes(b.api)
}

func (b *bootstrap) startCampaignService() {
	campaignRepository := repository.NewCampaignRepository(b.api.DB)
	campaignService := service.NewCampaignService(&campaignRepository)
	campaignController := controllers.NewCampaignController(&campaignService)

	initCampaignRoutes(b.api, campaignController)
}
