package db

import (
	"database/sql"
	"fmt"
	"recipevault/util"
)

// Get formatted database connection string
func getDBConnectionString(c *util.Config) string {
	return fmt.Sprintf(`
    host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_PASSWORD, c.DB_DATABASE)
}

// Connect to the database - panic if unsuccessfull
func ConnectDB(log *util.Logger, c *util.Config) *sql.DB {
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
