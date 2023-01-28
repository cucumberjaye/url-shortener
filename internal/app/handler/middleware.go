package handler

import (
	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"net/http"
	"time"
)

func (h *Handler) authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("authorization")
		if err == nil {
			id, err := h.AuthService.CheckToken(c.Value)
			if err == nil {
				h.AuthService.SetCurrentId(id)
				next.ServeHTTP(w, r)
				return
			}
		}
		token, err := h.AuthService.GenerateNewToken()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.ErrorLogger.Println(err.Error())
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "authorization",
			Value:   token,
			Expires: time.Now().Add(30 * 24 * time.Hour),
		})
		next.ServeHTTP(w, r)
	})
}
