package main

import (
	"net/http"
	"recipevault/api"
	"recipevault/db"
	"recipevault/util"

	_ "github.com/lib/pq"
)

// Run http server
func main() {
	log := util.NewLogger()
	conf := util.NewConfig()

	mux := http.NewServeMux()
	db := db.ConnectDB(log, conf)

	router := api.NewRouter(mux, log, db)

	server := http.Server{
		Addr:    ":" + conf.API_PORT,
		Handler: router,
	}

	log.Info("Starting server on port", conf.API_PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
