package api

import (
	"net/http"
	"recipevault/db"
	"recipevault/util"
)

// Start RecipeVault http server
func StartServer() {
	log := util.NewLogger()           // Custom logger
	conf := util.NewConfig()          // ENV config
	db := db.ConnectDB(log, conf)     // Database setup
	mux := http.NewServeMux()         // Mux server
	router := NewRouter(mux, log, db) // Router mappings

	server := http.Server{
		Addr:    ":" + conf.API_PORT,
		Handler: router,
	}

	log.Info("Starting server on port", conf.API_PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
