package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"recipehub/api/src/handler"
)

// Environment variables
var PORT = ":" + os.Getenv("API_PORT")

// Add routes and inject dependencies to handlers
func addRoutes(router *http.ServeMux, connection *sql.DB) {
	router.HandleFunc("GET /api", handler.Root)
}

// Creates a new router, adds routes and applies middleware
func newRouter() *http.ServeMux {
	router := http.NewServeMux()

	connection := NewDBConnection()

	addRoutes(router, connection)

	//handler = someMiddleware(handler)

	return router
}

// Start http server
func Start() {
	router := newRouter()

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Println("Starting server on port: %s", PORT)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}

}
