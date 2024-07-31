package types

import (
	"database/sql"
)

type Repository struct {
	connection *sql.DB
}
