package handler

import (
	"github.com/go-chi/chi"
)

type URLService interface {
	ShortingURL(fullURL string) (string, error)
	GetFullURL(shortURL string) (string, error)
}

type Handler struct {
	Service URLService
}

func NewHandler(service URLService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", h.shortener)
	r.Get("/{short}", h.getFullURL)

	return r
}
