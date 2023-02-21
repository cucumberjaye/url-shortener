package handler

import (
	mw "github.com/cucumberjaye/url-shortener/internal/app/middleware"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/models"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
)

func (h *Handler) getFullURL(w http.ResponseWriter, r *http.Request) {
	shortURL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	short := chi.URLParam(r, "short")
	short = baseURL(r) + short
	fullURL, err := h.Service.GetFullURL(short)
	if err != nil && err.Error() == "URL was deleted" {
		w.WriteHeader(http.StatusGone)
		return
	} else if err != nil {
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

	id, ok := r.Context().Value(mw.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		logger.ErrorLogger.Println("id must be string")
		return
	}

	fullURL := string(body)
	fullURL = strings.Trim(fullURL, "\n")
	shortURL, err := h.Service.ShortingURL(fullURL, baseURL(r), id)
	if err != nil && err.Error() == "url already exists" {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(shortURL))
		return
	} else if err != nil {
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

	id, ok := r.Context().Value(mw.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		logger.ErrorLogger.Println("id must be string")
		return
	}

	fullURL := input.URL
	shortURL, err := h.Service.ShortingURL(fullURL, baseURL(r), id)
	if err != nil && err.Error() == "url already exists" {
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, map[string]string{
			"result": shortURL,
		})
		return
	} else if err != nil {
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

	URL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	id, ok := r.Context().Value(mw.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		logger.ErrorLogger.Println("id must be string")
		return
	}

	out, err := h.Service.GetAllUserURL(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s  %s", r.Method, URL.String(), err.Error())
		return
	}

	if len(out) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, out)
}

func (h *Handler) checkDBConn(w http.ResponseWriter, r *http.Request) {
	URL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	err := h.Service.CheckDBConn()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s %s %s  ", r.Method, URL.String(), err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) batchShortener(w http.ResponseWriter, r *http.Request) {
	var input []models.BatchInputJSON
	var out = []models.BatchOutputJSON{}

	URL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	if err := render.DecodeJSON(r.Body, &input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, URL.String(), err.Error())
		return
	}

	id, ok := r.Context().Value(mw.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		logger.ErrorLogger.Println("id must be string")
		return
	}

	tmp, err := h.Service.BatchSetURL(input, baseURL(r), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, URL.String(), err.Error())
		return
	}

	for i := 0; i < len(tmp); i++ {
		out = append(out, models.BatchOutputJSON{
			CorrelationID: tmp[i].CorrelationID,
			ShortURL:      tmp[i].OriginalURL,
		})
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, out)

	for i := 0; i < len(out); i++ {
		logger.InfoLogger.Printf("%s  Full URL: %s has been added with short URL: %s", r.Method, input[i].OriginalURL, out[i].ShortURL)
	}

}

func (h *Handler) deleteUserURL(w http.ResponseWriter, r *http.Request) {
	var input []string

	URL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	id, ok := r.Context().Value(mw.UserID("user_id")).(string)
	if !ok {
		http.Error(w, "error on server", http.StatusInternalServerError)
		logger.ErrorLogger.Println("id must be string")
		return
	}

	if err := render.DecodeJSON(r.Body, &input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.WarningLogger.Printf("%s  %s: %s", r.Method, URL.String(), err.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)

	for i := range input {
		input[i] = r.URL.Scheme + r.Host + getURLPath + input[i]
	}
	h.Service.BatchDeleteURL(input, id)
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
