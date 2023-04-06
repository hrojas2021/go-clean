package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hugo.rojas/custom-api/internal/domain/models"
	"github.com/hugo.rojas/custom-api/internal/errors"
)

type payload struct {
	Name string `json:"name"`
}

func (h *Handle) SaveRoom(w http.ResponseWriter, req *http.Request) {
	var payload payload
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		h.resp.Failf(w, req, "could not parse body params; %w", errors.ErrInvalidPayload)
		return
	}

	if len(strings.TrimSpace(payload.Name)) == 0 {
		h.resp.Failf(w, req, "the params are invalid; %w", errors.ErrInvalidPayload)
		return
	}

	room := models.Room{Name: payload.Name}

	err = h.service.SaveRoom(req.Context(), &room)
	if err != nil {
		h.resp.Failf(w, req, "could not save the room; %w", err)
		return
	}
	h.resp.JSON(w, req, room)
}
