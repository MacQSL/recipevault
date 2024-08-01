package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"recipehub/api/src/service"
)

var PORT = ":" + os.Getenv("API_PORT")

// Start http server
func Start() {
	mux := http.NewServeMux()
	db := SetupDatabase()

	router := NewRouter(mux, db)

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
