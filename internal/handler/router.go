package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/hisshihi/url-shortener/internal/service"
)

type URLService interface {
	CreateShortURL(ctx context.Context, rawUrl string) (string, error)
	SelectByAlias(ctx context.Context, alias string) (string, error)
}

type Handler struct {
	svc URLService
}

func NewUrlHandler(svc URLService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			return
		}
	})

	mux.HandleFunc("POST /api/v1/urls", h.create)
	mux.HandleFunc("GET /api/v1/redirect", h.getByAlias)

	return mux
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
