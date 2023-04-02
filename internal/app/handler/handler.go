// Модуль для приема запросов к серверу и возвращения ответов.
package handler

import (
	mw "github.com/cucumberjaye/url-shortener/internal/app/middleware"
	"github.com/cucumberjaye/url-shortener/models"

	"github.com/go-chi/chi"
)

const (
	getURLPath = "/"
)

// LogsInfoService сервис для логирования запросов.
type LogsInfoService interface {
	GetRequestCount(shortURL string) (int, error)
}

// URLService основной сервис, обрабатывающий запросы.
type URLService interface {
	ShortingURL(fullURL, baseURL string, id string) (string, error)
	GetFullURL(shortURL string) (string, error)
	GetAllUserURL(id string) ([]models.URLs, error)
	CheckDBConn() error
	BatchSetURL(data []models.BatchInputJSON, baseURL string, id string) ([]models.BatchInputJSON, error)
}

// Handler хранит обЪекты сервисов для их испльзования.
type Handler struct {
	Service       URLService
	LoggerService LogsInfoService
	Ch            chan models.DeleteData
}

// NewHandler создает объект Handler
func NewHandler(service URLService, logsService LogsInfoService, ch chan models.DeleteData) *Handler {
	return &Handler{
		Service:       service,
		LoggerService: logsService,
		Ch:            ch,
	}
}

// InitRoutes возвращает роутер с эндпоинтами и подключенными middleware
func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.With(mw.GzipCompress, mw.Authentication).Get("/{short}", h.GetFullURL)

	r.Get("/ping", h.CheckDBConn)
	r.With(mw.Authentication).Group(func(r chi.Router) {
		r.Use(mw.GzipDecompress)
		r.Post("/", h.Shortener)
		r.Route("/api", func(r chi.Router) {
			r.With(mw.GzipDecompress).Route("/shorten", func(r chi.Router) {
				r.Post("/", h.ShortenerJSON)
				r.Post("/batch", h.BatchShortener)
			})

			r.Route("/user", func(r chi.Router) {
				r.Get("/urls", h.GetUserURL)
				r.Delete("/urls", h.DeleteUserURL)
			})
		})
	})

	return r
}
