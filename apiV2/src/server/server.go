package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"recipehub/api/src/handler"
)

var PORT = ":" + os.Getenv("API_PORT")

// Add routes and inject dependencies to handlers
func addRoutes(router *http.ServeMux, db *sql.DB) {
	router.HandleFunc("GET /api", handler.Root)
}

// Creates a new router, adds routes and applies middleware
func newRouter(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	addRoutes(router, db)

	//router = someMiddleware(handler)

	return router
}

// Start http server
func Start() {
	db := SetupDatabase()
	router := newRouter(db)

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Println("Starting server on port", PORT)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
