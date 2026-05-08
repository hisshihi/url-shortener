package handler

import (
	"context"
	"log"
	"net/http"
)

type URLService interface {
	CreateShortURL(ctx context.Context, rawUrl string) (string, error)
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
		write, err := w.Write([]byte("OK"))
		if err != nil {
			return
		}
		log.Println(write)
	})

	mux.HandleFunc("POST /api/v1/urls", h.create)

	return mux
}
