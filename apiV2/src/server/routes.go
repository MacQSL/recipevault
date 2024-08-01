package server

import (
	"database/sql"
	"net/http"
	"recipehub/api/src/handler"
	"recipehub/api/src/middleware"
)

// Add routes and inject dependencies to handlers
func addRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("GET /health", handler.Health)

	mux.HandleFunc("/", http.NotFound)
}

// Creates a new router, adds routes and applies middleware
func NewRouter(mux *http.ServeMux, db *sql.DB) http.Handler {
	// Prepend `/api/` to all routes
	mux.Handle("/api/", http.StripPrefix("/api", mux))

	addRoutes(mux, db)

	var router http.Handler = mux

	// Apply middleware to all routes
	router = middleware.Logger(router)

	return router
}
