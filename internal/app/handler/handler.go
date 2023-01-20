package handler

import (
	"github.com/go-chi/chi"
)

const (
	protocol   = "http"
	getURLPath = "/"
)

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

	r.With(h.gzipCompress).Get("/{short}", h.getFullURL)

	r.Group(func(r chi.Router) {
		r.Use(h.gzipDecompress)
		r.Post("/", h.shortener)
		r.Route("/api", func(r chi.Router) {
			r.With(h.gzipDecompress).Post("/shorten", h.JSONShortener)
		})
	})

	return r
}
