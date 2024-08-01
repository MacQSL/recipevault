package server

import (
	"net/http"
	"recipehub/api/src/util"
	"time"
)

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// Logs incoming requests ie: 200 GET /api/path 10
func LoggerMiddleware(log *util.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wrapped := &wrappedResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			next.ServeHTTP(wrapped, r)

			log.Info(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
		}
		return http.HandlerFunc(fn)
	}
}
