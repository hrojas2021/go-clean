package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hugo.rojas/custom-api/internal/domain/models"
)

func (h *Handle) ListUsers(w http.ResponseWriter, req *http.Request) {

	users, err := h.service.ListUser(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func (h *Handle) Login(w http.ResponseWriter, req *http.Request) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if user.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("username is required"))
		return
	}

	if user.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("password is required"))
		return
	}

	err = h.service.Login(req.Context(), user)
	_ = err
}
