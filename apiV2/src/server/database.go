package server

import (
	"database/sql"
	"recipehub/api/src/util"
)

// Connect to the database - panic if unsuccessfull
func ConnectDB(log util.ILogger, connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Unable to ping database")
	}

	log.Info("DB connected")

	return db
}
