package handler

import (
	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"net/url"
)

func (h *Handler) getFullURL(w http.ResponseWriter, r *http.Request) {
	shortURL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	short := chi.URLParam(r, "short")
	fullURL, err := h.Service.GetFullURL(short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, shortURL.String(), err.Error())
		return
	}
	w.Header().Set("Location", fullURL)
	w.WriteHeader(307)

	requestCount, err := h.LoggerService.GetRequestCount(short)
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

	fullURL := string(body)
	shortURL, err := h.Service.ShortingURL(fullURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s:  fullURL: %s %s", r.Method, URL.String(), fullURL, err.Error())
		return
	}
	
	shortURL = baseURL(r) + shortURL

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))

	logger.InfoLogger.Printf("%s  Full URL: %s has been added with short URL: %s", r.Method, fullURL, shortURL)
}

type JSONInput struct {
	URL string `json:"url"`
}

func (h *Handler) JSONShortener(w http.ResponseWriter, r *http.Request) {
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

	fullURL := input.URL
	shortURL, err := h.Service.ShortingURL(fullURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s:  fullURL: %s %s", r.Method, URL.String(), fullURL, err.Error())
		return
	}
	shortURL = baseURL(r) + shortURL

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{
		"result": shortURL,
	})

	logger.InfoLogger.Printf("%s  Full URL: %s has been added with short URL: %s", r.Method, fullURL, shortURL)
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
