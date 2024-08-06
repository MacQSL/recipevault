package middleware

import (
	"net/http"
	"recipevault/util"
	"time"
)

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int // to not overwrite 'status' in original struct
}

// Wrap the existing implementation of WriteHeader to intercept
// the status code when it is first passed to WriteHeader
func (w *wrappedResponseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.statusCode = status
}

// Logs incoming requests ie: 200 GET /api/path 10
func LoggerMiddleware(log *util.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wrapped := &wrappedResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			next.ServeHTTP(wrapped, r)

			log.Info(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
		})
	}
}
