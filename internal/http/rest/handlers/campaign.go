package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/uptrace/bunrouter"
)

func (h *Handle) GetCampaign(w http.ResponseWriter, req *http.Request) {
	params := bunrouter.ParamsFromContext(req.Context())
	campaignID := params.ByName("id")

	w.Header().Set("Content-Type", "application/json")

	campID, err := uuid.Parse(campaignID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	campaign, err := h.service.GetCampaign(req.Context(), campID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	if err := json.NewEncoder(w).Encode(campaign); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
