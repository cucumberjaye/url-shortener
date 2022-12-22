package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"net/http"
)

func (h *Handler) getFullURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "short")
	fullURL, err := h.Service.Shortener.GetFullURL(shortURL)
	if err != nil {
		fmt.Println(shortURL)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", fullURL)
	w.WriteHeader(307)
}

func (h *Handler) shortener(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(body) == 0 {
		http.Error(w, "body is empty", http.StatusBadRequest)
		return
	}
	shortURL, err := h.Service.Shortener.ShortingURL(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(shortURL))
}

/*func (h *Handler) Shortener(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/" {
			http.Error(w, "expect /<id>", http.StatusBadRequest)
			return
		}
		shortURL := r.URL.Path[1:]
		fullURL, err := h.Service.Shortener.GetFullURL(shortURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", fullURL)
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
		shortURL, err := h.Service.Shortener.ShortingURL(string(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

		w.Write([]byte(shortURL))
	}
}*/
