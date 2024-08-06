package api

import (
	"database/sql"
	"net/http"
	"recipevault/api/handler"
	"recipevault/api/middleware"
	"recipevault/util"
)

// Add routes and inject dependencies to handlers
func addRoutes(mux *http.ServeMux, log *util.Logger, db *sql.DB) {
	// Cookbook ID related endpoints require user authorization
	cookbookAuth := middleware.CookbookAuthMiddleware(log, db)

	mux.Handle("GET /health", handler.GetHealth())
	// Get all user cookbooks
	mux.Handle("GET /cookbooks", handler.GetUserCookbooks(log, db))
	// Get cookbook by ID
	mux.Handle("GET /cookbooks/{cookbookID}", cookbookAuth(handler.GetUserCookbook(log, db)))
	// Get all cookbook recipes
	mux.Handle("GET /cookbooks/{cookbookID}/recipes", cookbookAuth(handler.GetUserCookbook(log, db)))

	mux.Handle("/", http.NotFoundHandler())
}

// Creates a new router, adds routes and applies middleware
func NewRouter(mux *http.ServeMux, log *util.Logger, db *sql.DB) http.Handler {
	mux.Handle("/api/", http.StripPrefix("/api", mux)) // prepend /api/

	addRoutes(mux, log, db)

	var router http.Handler = mux

	// Generate middleware
	logMiddleware := middleware.LoggerMiddleware(log)
	authMiddleware := middleware.AuthMiddleware(log)

	// Apply middleware to ALL routes
	router = logMiddleware(router)
	router = middleware.HeadersMiddleware(router)
	router = authMiddleware(router)

	return router
}
