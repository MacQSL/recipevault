package server

import (
	"database/sql"
	"net/http"
	"recipevault/api/src/handler"
	"recipevault/api/src/util"
)

// Add routes and inject dependencies to handlers
func addRoutes(mux *http.ServeMux, log util.ILogger, db *sql.DB) {
	mux.HandleFunc("GET /health", handler.GetHealth)
	mux.HandleFunc("GET /cookbooks", handler.GetUserCookbooks(log, db))
	mux.HandleFunc("GET /cookbooks/{cookbookID}", handler.GetCookbook(log, db))

	mux.HandleFunc("/", http.NotFound)
}

// Creates a new router, adds routes and applies middleware
func NewRouter(mux *http.ServeMux, log util.ILogger, db *sql.DB) http.Handler {
	mux.Handle("/api/", http.StripPrefix("/api", mux)) // prepend /api/

	addRoutes(mux, log, db)

	var router http.Handler = mux

	// Generate middleware
	logMiddleware := LoggerMiddleware(log)
	authMiddleware := AuthMiddleware(log)

	// Apply middleware to ALL routes
	router = logMiddleware(router)
	router = HeadersMiddleware(router)
	router = authMiddleware(router)

	return router
}
