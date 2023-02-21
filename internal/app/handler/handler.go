package handler

import (
	mw "github.com/cucumberjaye/url-shortener/internal/app/middleware"
	"github.com/cucumberjaye/url-shortener/models"

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
	ShortingURL(fullURL, baseURL string, id string) (string, error)
	GetFullURL(shortURL string) (string, error)
	GetAllUserURL(id string) ([]models.URLs, error)
	CheckDBConn() error
	BatchSetURL(data []models.BatchInputJSON, baseURL string, id string) ([]models.BatchInputJSON, error)
	BatchDeleteURL(data []string, id string)
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

	r.With(mw.GzipCompress, mw.Authentication).Get("/{short}", h.getFullURL)

	r.Get("/ping", h.checkDBConn)
	r.With(mw.Authentication).Group(func(r chi.Router) {
		r.Use(mw.GzipDecompress)
		r.Post("/", h.shortener)
		r.Route("/api", func(r chi.Router) {
			r.With(mw.GzipDecompress).Route("/shorten", func(r chi.Router) {
				r.Post("/", h.shortenerJSON)
				r.Post("/batch", h.batchShortener)
			})

			r.Route("/user", func(r chi.Router) {
				r.Get("/urls", h.getUserURL)
				r.Delete("/urls", h.deleteUserURL)
			})
		})
	})

	return r
}
