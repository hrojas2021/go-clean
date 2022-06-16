package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hugo.rojas/custom-api/internal/domain/models"
	service "github.com/hugo.rojas/custom-api/internal/service/campaign"
	"github.com/julienschmidt/httprouter"
)

type CampaignController struct {
	campaignService service.CampaignService
}

func NewCampaignController(campaignService *service.CampaignService) *CampaignController {
	return &CampaignController{
		campaignService: *campaignService,
	}
}

func (c *CampaignController) GetCampaign(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	campaignID := params.ByName("campaignID")
	var request models.GetCampaignRequest
	request.ID = campaignID

	w.Header().Set("Content-Type", "application/json")
	campaign, err := c.campaignService.GetCampaign(request)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		w.Write([]byte(err.Err.Error()))
	}
	if err := json.NewEncoder(w).Encode(campaign); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}
