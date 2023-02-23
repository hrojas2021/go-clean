package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hugo.rojas/custom-api/internal/domain/models"
)

func (h *Handle) SaveRoom(w http.ResponseWriter, req *http.Request) {
	var room models.Room
	err := json.NewDecoder(req.Body).Decode(&room)
	if err != nil {
		h.resp.Failf(w, req, "could not parse body params; %w", err)
		return
	}

	err = h.service.SaveRoom(req.Context(), &room)
	if err != nil {
		h.resp.Fail(w, req, err)
		return
	}

	h.resp.JSON(w, req, room)
}
