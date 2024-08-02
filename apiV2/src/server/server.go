package server

import (
	"net/http"
	"recipehub/api/src/util"
)

// Run http server
func Run() {
	log := util.NewLogger()
	conf := util.NewConfig()

	mux := http.NewServeMux()
	db := ConnectDB(log, conf)

	router := NewRouter(mux, log, db)

	server := http.Server{
		Addr:    ":" + conf.API_PORT,
		Handler: router,
	}

	log.Info("Starting server on port", conf.API_PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
