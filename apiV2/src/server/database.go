package server

import (
	"database/sql"
	"fmt"
	"recipehub/api/src/util"
)

// Get formatted database connection string
func getDBConnectionString(c *util.Config) string {
	return fmt.Sprintf(`
    host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBDatabase)
}

// Connect to the database - panic if unsuccessfull
func ConnectDB(log util.ILogger, c *util.Config) *sql.DB {
	db, err := sql.Open("postgres", getDBConnectionString(c))

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
