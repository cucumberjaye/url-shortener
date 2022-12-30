package handler

import (
	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"github.com/go-chi/chi"
	"io"
	"net/http"
)

func (h *Handler) getFullURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "short")
	fullURL, err := h.Service.GetFullURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, protocol+"://"+r.Host+r.URL.Path, err.Error())
		return
	}
	w.Header().Set("Location", fullURL)
	w.WriteHeader(307)

	requestCount, err := h.LoggerService.GetRequestCount(shortURL)
	if err != nil {
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, protocol+"://"+r.Host+r.URL.Path, err.Error())
		return
	}
	logger.InfoLogger.Printf("%s  URL: %s has been used, total uses: %d", r.Method, protocol+"://"+r.Host+r.URL.Path, requestCount)

}

func (h *Handler) shortener(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, protocol+"://"+r.Host+r.URL.Path, err.Error())
		return
	}
	if len(body) == 0 {
		http.Error(w, "body is empty", http.StatusBadRequest)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, protocol+"://"+r.Host+r.URL.Path, "body is empty")
		return
	}
	fullURL := string(body)
	shortURL, err := h.Service.ShortingURL(fullURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s:  fullURL: %s %s", r.Method, protocol+"://"+r.Host+r.URL.Path, fullURL, err.Error())
		return
	}
	shortURL = "http://" + r.Host + r.URL.Path + shortURL

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))

	logger.InfoLogger.Printf("%s  Full URL: %s has been added with short URL: %s", r.Method, fullURL, shortURL)
}
