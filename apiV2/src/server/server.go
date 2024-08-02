package server

import (
	"log"
	"net/http"
	"recipehub/api/src/util"
)

// Run http server
func Run() {
	log := util.NewLogger(&log.Logger{})
	conf := util.NewConfig()

	mux := http.NewServeMux()
	db := ConnectDB(log, conf.GetDBConnectionString())
	router := NewRouter(mux, log, db)

	server := http.Server{
		Addr:    ":" + conf.DBPort,
		Handler: router,
	}

	log.Info("Starting server on port", conf.DBPort)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
