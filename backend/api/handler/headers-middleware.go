package handler

import (
	"net/http"
)

// Apply headers
func HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// JSON
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
