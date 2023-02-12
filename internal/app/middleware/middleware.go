package middleware

import (
	"compress/gzip"
	"context"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cucumberjaye/url-shortener/pkg/logger"
	"github.com/cucumberjaye/url-shortener/pkg/token"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}

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

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("authorization")
		if err == nil {
			id, err := token.CheckToken(c.Value)
			if err == nil {
				userId := "user_id"
				ctx := context.WithValue(r.Context(), userId, id)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		id := uuid.New().String()
		userId := "user_id"
		ctx := context.WithValue(r.Context(), userId, id)
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
