package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hugo.rojas/custom-api/internal/domain/models"
	"github.com/hugo.rojas/custom-api/internal/errors"
)

func (h *Handle) ListUsers(w http.ResponseWriter, req *http.Request) {
	users, err := h.service.ListUser(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.resp.JSON(w, req, map[string]interface{}{
		"users": users,
	})
}

func (h *Handle) Login(w http.ResponseWriter, req *http.Request) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		h.resp.Failf(w, req, "could not parse body params; %w", err)
		return
	}

	if user.Username == "" {
		h.resp.Fail(w, req, errors.ErrInvalidPayload)
		return
	}

	if user.Username == "" {
		h.resp.Fail(w, req, errors.ErrInvalidPayload)
		return
	}

	token, err := h.service.Login(req.Context(), user)
	if err != nil {
		h.resp.Fail(w, req, err)
		return
	}

	h.resp.JSON(w, req, token)
}
