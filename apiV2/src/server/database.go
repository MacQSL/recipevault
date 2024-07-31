package server

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	host     = os.Getenv("API_HOST")
	port     = os.Getenv("API_PORT")
	user     = os.Getenv("DB_USER_API")
	password = os.Getenv("DB_USER_API_PASSWORD")
	database = os.Getenv("DB_DATABASE")
)

// Create a new database connection
func NewDBConnection() *sql.DB {
	psqlInfo := fmt.Sprintf(`
    host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, database)

	connection, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	err = connection.Ping()

	if err != nil {
		panic(err)
	}

	return connection
}