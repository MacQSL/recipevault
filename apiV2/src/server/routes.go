package server

import (
	"database/sql"
	"net/http"
	"recipehub/api/src/handler"
	"recipehub/api/src/util"
)

// Add routes and inject dependencies to handlers
func addRoutes(mux *http.ServeMux, log *util.Logger, db *sql.DB) {
	mux.HandleFunc("GET /health", handler.Health)

	mux.HandleFunc("/", http.NotFound)
}

// Creates a new router, adds routes and applies middleware
func NewRouter(mux *http.ServeMux, log *util.Logger, db *sql.DB) http.Handler {
	// Prepend `/api/` to all routes
	mux.Handle("/api/", http.StripPrefix("/api", mux))

	addRoutes(mux, log, db)

	var router http.Handler = mux

	// Generate middleware
	logMiddleware := LoggerMiddleware(log) // inject custom logger

	// Apply middleware to ALL routes
	router = logMiddleware(router)

	return router
}
