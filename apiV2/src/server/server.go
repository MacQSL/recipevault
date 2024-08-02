package server

import (
	"log"
	"net/http"
	"os"
	"recipehub/api/src/util"
)

var PORT = ":" + os.Getenv("API_PORT")

// Start http server
func Run() {
	mux := http.NewServeMux()
	db := ConnectDB()
	log := util.NewLogger(&log.Logger{})
	log.SetLogLevel("DEBUG")

	router := NewRouter(mux, log, db)

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Info("Starting server on port", PORT)

	err := server.ListenAndServe()

	if err != nil {
		log.Error("Error starting server:", err)
		os.Exit(1)
	}
}
