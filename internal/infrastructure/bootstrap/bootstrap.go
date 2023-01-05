package bootstrap

import (
	"github.com/hugo.rojas/custom-api/internal/http/rest"
	"github.com/hugo.rojas/custom-api/internal/http/rest/handlers"
	"github.com/hugo.rojas/custom-api/internal/infrastructure/api"
	repository "github.com/hugo.rojas/custom-api/internal/repository/campaign"
	service "github.com/hugo.rojas/custom-api/internal/service/campaign"
)

type Bootstraper interface {
	startNoopService()
	startCampaignService()
}

type bootstrap struct {
	api *api.API
}

func NewBootstrap(a *api.API) Bootstraper {
	return &bootstrap{
		api: a,
	}
}

func InitServices(b Bootstraper) {
	b.startNoopService()
	b.startCampaignService()
}

func (b *bootstrap) startNoopService() {
	rest.InitNoopRoutes(b.api)
}

func (b *bootstrap) startCampaignService() {
	campaignRepository := repository.NewCampaignRepository(b.api.DB)
	campaignService := service.NewCampaignService(&campaignRepository)
	campaignController := handlers.NewCampaignHandler(&campaignService)

	rest.InitCampaignRoutes(b.api, campaignController)
}
