package handler

import (
	"fmt"
	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/models"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (h *Handler) getFullURL(w http.ResponseWriter, r *http.Request) {
	shortURL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	id := h.AuthService.GetCurrentId()

	short := chi.URLParam(r, "short")
	short = baseURL(r) + short
	fullURL, err := h.Service.GetFullURL(short, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, shortURL.String(), err.Error())
		return
	}
	w.Header().Set("Location", fullURL)
	w.WriteHeader(307)

	requestCount, err := h.LoggerService.GetRequestCount(short, id)
	if err != nil {
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, shortURL.String(), err.Error())
		return
	}
	logger.InfoLogger.Printf("%s  URL: %s has been used, total uses: %d", r.Method, shortURL.String(), requestCount)
}

func (h *Handler) shortener(w http.ResponseWriter, r *http.Request) {
	URL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, URL.String(), err.Error())
		return
	}

	if len(body) == 0 {
		http.Error(w, "body is empty", http.StatusBadRequest)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, URL.String(), "body is empty")
		return
	}

	id := h.AuthService.GetCurrentId()

	fullURL := string(body)
	fullURL = strings.Trim(fullURL, "\n")
	fmt.Println(fullURL)
	shortURL, err := h.Service.ShortingURL(fullURL, baseURL(r), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s:  fullURL: %s %s", r.Method, URL.String(), fullURL, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))

	logger.InfoLogger.Printf("%s  Full URL: %s has been added with short URL: %s", r.Method, fullURL, shortURL)
}

type JSONInput struct {
	URL string `json:"url"`
}

func (h *Handler) shortenerJSON(w http.ResponseWriter, r *http.Request) {
	var input = &JSONInput{}

	URL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	if err := render.DecodeJSON(r.Body, input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, URL.String(), err.Error())
		return
	}

	if len(input.URL) == 0 {
		http.Error(w, "body is empty", http.StatusBadRequest)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, URL.String(), "body is empty")
		return
	}

	id := h.AuthService.GetCurrentId()

	fullURL := input.URL
	shortURL, err := h.Service.ShortingURL(fullURL, baseURL(r), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s:  fullURL: %s %s", r.Method, URL.String(), fullURL, err.Error())
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{
		"result": shortURL,
	})

	logger.InfoLogger.Printf("%s  Full URL: %s has been added with short URL: %s", r.Method, fullURL, shortURL)
}

func (h *Handler) getUserURL(w http.ResponseWriter, r *http.Request) {
	var out []models.URLs

	id := h.AuthService.GetCurrentId()
	out = h.Service.GetAllUserURL(id)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, out)
}

func baseURL(r *http.Request) string {
	if configs.BaseURL != "" {
		return configs.BaseURL + "/"
	}

	result := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   getURLPath,
	}

	return result.String()
}
