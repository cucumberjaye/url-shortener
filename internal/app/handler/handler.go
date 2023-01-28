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
	GetRequestCount(shortURL string, id int) (int, error)
}

type URLService interface {
	ShortingURL(fullURL, baseURL string, id int) (string, error)
	GetFullURL(shortURL string, id int) (string, error)
	GetAllUserURL(id int) []models.URLs
}

type AuthService interface {
	GenerateNewToken() (string, error)
	CheckToken(token string) (int, error)
	SetCurrentID(id int)
	GetCurrentID() int
}

type Handler struct {
	Service       URLService
	AuthService   AuthService
	LoggerService LogsInfoService
}

func NewHandler(service URLService, logsService LogsInfoService, authService AuthService) *Handler {
	return &Handler{
		Service:       service,
		LoggerService: logsService,
		AuthService:   authService,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(h.authentication)
	r.With(mw.GzipCompress).Get("/{short}", h.getFullURL)

	r.Group(func(r chi.Router) {
		r.Use(mw.GzipDecompress)
		r.Post("/", h.shortener)
		r.Route("/api", func(r chi.Router) {
			r.With(mw.GzipDecompress).Post("/shorten", h.shortenerJSON)

			r.Route("/user", func(r chi.Router) {
				r.Get("/urls", h.getUserURL)
			})
		})
	})

	return r
}
