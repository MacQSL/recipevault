package db

import (
	"database/sql"
	"recipevault/util"
)

// Connect to the database - panic if unsuccessfull
func ConnectDB(log *util.Logger, c *util.Config) *sql.DB {
	db, err := sql.Open("postgres", c.GetDBConnectionURL())

	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Unable to ping database")
	}

	log.Info("DB connected!")

	return db
}
