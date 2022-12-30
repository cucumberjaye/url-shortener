package handler

import (
	"github.com/go-chi/chi"
)

const protocol = "http"

type LogsInfoService interface {
	GetRequestCount(shortURL string) (int, error)
}

type URLService interface {
	ShortingURL(fullURL string) (string, error)
	GetFullURL(shortURL string) (string, error)
}

type Handler struct {
	Service       URLService
	LoggerService LogsInfoService
}

func NewHandler(service URLService, logsService LogsInfoService) *Handler {
	return &Handler{
		Service:       service,
		LoggerService: logsService,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", h.shortener)
	r.Get("/{short}", h.getFullURL)

	return r
}
