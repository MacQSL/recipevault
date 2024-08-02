package server

import (
	"log"
	"net/http"
	"os"
	"recipehub/api/src/util"
)

// Run http server
func Run() {
	// API dependencies
	log := util.NewLogger(&log.Logger{})
	mux := http.NewServeMux()
	db := ConnectDB(log)
	router := NewRouter(mux, log, db)

	port, ok := os.LookupEnv("API_PORT")

	if !ok {
		log.Error("Missing API_PORT")
		os.Exit(1)
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Info("Starting server on port", port)

	err := server.ListenAndServe()

	if err != nil {
		log.Error("Error starting server:", err)
		os.Exit(1)
	}
}
