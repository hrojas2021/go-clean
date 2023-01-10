package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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
