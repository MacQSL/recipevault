package api

import (
	"net/http"
	"recipevault/db"
	"recipevault/util"
)

// Start RecipeVault http server
func StartServer() {
	log := util.NewLogger()
	conf := util.NewConfig()
	db := db.ConnectDB(log, conf)
	mux := http.NewServeMux()
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
