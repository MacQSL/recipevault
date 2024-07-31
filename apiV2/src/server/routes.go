package server

import (
	"database/sql"
	"net/http"
	"recipehub/api/src/handler"
)

// Add routes and inject dependencies to handlers
func addRoutes(router *http.ServeMux, db *sql.DB) {
	router.HandleFunc("GET /api", handler.Root)
}

// Creates a new router, adds routes and applies middleware
func NewRouter(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	addRoutes(router, db)

	//router = someMiddleware(handler)

	return router
}
