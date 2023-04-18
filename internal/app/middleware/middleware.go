package middleware

import (
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"github.com/cucumberjaye/url-shortener/pkg/token"
)

// для передачи зашифрованного тела далее
type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// метод структуры gzipWriter для записи
func (w gzipWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}

// GzipCompress сжимает тело ответа в gzip
func GzipCompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

// GzipDecompress возвращает сжатые данные в нормальный вид
func GzipDecompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reader io.Reader

		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			reader = gz
			defer gz.Close()
		} else {
			reader = r.Body
		}

		r.Body = io.NopCloser(reader)

		next.ServeHTTP(w, r)
	})
}

// для передачи id в контексте
type UserID string

// Authentication проверяет авторизован ли пользователь
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("authorization")
		if err == nil {
			var id string
			id, err = token.CheckToken(c.Value)
			if err == nil {
				ctx := context.WithValue(r.Context(), UserID("user_id"), id)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		id := uuid.New().String()
		ctx := context.WithValue(r.Context(), UserID("user_id"), id)
		authToken, err := token.GenerateNewToken(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.ErrorLogger.Println(err.Error())
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "authorization",
			Value:   authToken,
			Expires: time.Now().Add(30 * 24 * time.Hour),
			Path:    "/",
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
