package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) getByAlias(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Alias string `json:"alias"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Ошибка при передачи данных", http.StatusBadRequest)
		return
	}

	alias, err := h.svc.SelectByAlias(r.Context(), request.Alias)
	if err != nil {
		respondError(w, err)
	}

	var response struct {
		URL string `json:"url"`
	}

	response.URL = alias

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMovedPermanently)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		respondError(w, err)
		return
	}
}
