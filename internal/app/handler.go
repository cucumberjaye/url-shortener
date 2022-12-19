package app

import (
	"io"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Shortener(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/" {
			http.Error(w, "expect /<id>", http.StatusBadRequest)
			return
		}
		shortUrl := r.URL.Path[1:]
		fullUrl, err := h.Service.GetFullURL(shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", fullUrl)
		w.WriteHeader(307)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(body) == 0 {
			http.Error(w, "body is empty", http.StatusBadRequest)
		}
		shortUrl, err := h.Service.ShortingURL(string(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

		w.Write([]byte(shortUrl))
	}
}
