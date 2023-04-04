package handler

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	mw "github.com/cucumberjaye/url-shortener/internal/app/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/cucumberjaye/url-shortener/configs"
	"github.com/cucumberjaye/url-shortener/models"
	"github.com/cucumberjaye/url-shortener/pkg/logger"
)

// getFullURL перенаправляет на полную ссылку, id короткой ссылки находится после символа /.
func (h *Handler) getFullURL(w http.ResponseWriter, r *http.Request) {
	shortURL := url.URL{
		Scheme: configs.Scheme,
		Host:   r.Host,
		Path:   r.URL.Path,
	}

	short := baseURL(r) + chi.URLParam(r, "short")
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

// shortener приниает ссылку в формате text, возвращает короткую ссылку.
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

	fullURL := strings.Trim(string(body), "\n")
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

// JSON структура, которую принимает ShortenerJSON
type JSONInput struct {
	URL string `json:"url"`
}

// shortenerJSON приниает ссылку в формате JSON, возвращает короткую ссылку.
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

// getUserURL возвращает все сокращенные ссылки пользователя.
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

// checkDBConn проверяет работоспособность хранилища (postgreSQL или файла).
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

// batchShortener приниает массив ссылок в формате JSON (с двумя полями в структуре), возвращает массив коротких ссылок.
func (h *Handler) batchShortener(w http.ResponseWriter, r *http.Request) {
	var input []models.BatchInputJSON

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
	length := len(tmp)
	out := make([]models.BatchOutputJSON, length)

	for i := 0; i < length; i++ {
		out[i] = models.BatchOutputJSON{
			CorrelationID: tmp[i].CorrelationID,
			ShortURL:      tmp[i].OriginalURL,
		}
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, out)

	for i := 0; i < len(out); i++ {
		logger.InfoLogger.Printf("%s  Full URL: %s has been added with short URL: %s", r.Method, input[i].OriginalURL, out[i].ShortURL)
	}

}

// deleteUserURL принимает массив коротких ссылок пользователя и удаляет их.
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
	w.Write([]byte{})

	for i := range input {
		input[i] = baseURL(r) + input[i]
		h.Ch <- models.DeleteData{
			ID:       id,
			ShortURL: input[i],
		}
	}
}

// baseURL формирует корроткую ссылку
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
