package gzip

import (
	"compress/gzip"
	"github.com/MorZLE/url-shortener/internal/app/logger"
	"net/http"
	"strings"
)

func GzipMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
			if err != nil {
				logger.Error("Error creating gzip writer", err)
				return
			}
			defer gz.Close()
			w.Header().Set("Content-Encoding", "gzip")
		}

		if r.Header.Get(`Content-Encoding`) == `gzip` {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			r.Body = gz
			defer gz.Close()
		}

		// передаём управление хендлеру
		h.ServeHTTP(w, r)
	}
}
