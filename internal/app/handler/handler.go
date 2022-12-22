package handler

import (
	"github.com/cucumberjaye/url-shortener/internal/app/service"
	"github.com/go-chi/chi"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", h.shortener)
	r.Get("/{short}", h.getFullURL)

	return r
}
