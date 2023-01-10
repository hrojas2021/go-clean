package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
)

func (h *Handle) SaveRoom(w http.ResponseWriter, req *http.Request) {
	var room models.Room
	err := json.NewDecoder(req.Body).Decode(&room)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var r entities.Room
	r, err = h.service.SaveRoom(req.Context(), room)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
