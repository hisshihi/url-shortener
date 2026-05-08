package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hisshihi/url-shortener/internal/service"
)

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	url, err := h.svc.CreateShortURL(r.Context(), request.URL)
	if err != nil {
		respondError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func respondError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidURL):
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
