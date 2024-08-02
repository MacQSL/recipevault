package server

import (
	"database/sql"
	"fmt"
	"os"
	"recipehub/api/src/util"
)

// Format the database connection string
func connectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER_API")
	password := os.Getenv("DB_USER_API_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	return fmt.Sprintf(`
    host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, database)
}

// Connect to the database - panic if unsuccessfull
func ConnectDB(log util.ILogger) *sql.DB {
	db, err := sql.Open("postgres", connectionString())

	if err != nil {
		log.Error("Unable to connect to database.")
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Error("Unable to ping database.")
		panic(err)
	}

	log.Info("DB connected")

	return db
}
