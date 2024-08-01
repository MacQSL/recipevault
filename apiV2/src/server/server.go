package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"recipehub/api/src/util"
)

var PORT = ":" + os.Getenv("API_PORT")

// Start http server
func Run() {
	log := util.NewLogger(&log.Logger{})
	util.UpdateLogLevel(log, "DEBUG")

	mux := http.NewServeMux()

	db := SetupDatabase()

	router := NewRouter(mux, log, db)

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Info("Starting server on port", PORT)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
