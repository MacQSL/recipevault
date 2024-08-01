package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var PORT = ":" + os.Getenv("API_PORT")

// Start http server
func Start() {
	log := NewLogger(&log.Logger{})
	mux := http.NewServeMux()
	db := SetupDatabase()
	router := NewRouter(mux, log, db)

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.setLogLevel("DEBUG")
	log.Info("Starting server on port", PORT)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
