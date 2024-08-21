package repository

import (
	"database/sql"
	"errors"
)

// Models must implement Ptrs() method to return pointers to struct fields
type Row[T any] interface {
	Ptrs() []any
	*T
}

// Get single entity from database and scan into struct
func GetRow[T any, PT Row[T]](db *sql.DB, query string, args ...any) (T, error) {
	row := db.QueryRow(query, args...)

	var t T
	ptr := PT(&t)

	// Scan row into struct
	if err := row.Scan(ptr.Ptrs()...); err != nil {
		if err == sql.ErrNoRows {
			return t, errors.New("No row found")
		}

		// Return error if row scan fails
		return t, err
	}

	return t, nil
}

// Get multiple entities from database and scan into structs
func FindRows[T any, PT Row[T]](db *sql.DB, query string, args ...any) ([]T, error) {
	rows, err := db.Query(query, args...)

	// Handle database errors
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Return nil slice if no rows found ie: []
			return nil, nil
		}

		// Return error if query fails
		return nil, err
	}

	defer rows.Close()

	var result []T

	// Iterate over rows and scan into struct
	for rows.Next() {
		var t T
		ptr := PT(&t)

		if err := rows.Scan(ptr.Ptrs()...); err != nil {
			return nil, err
		}

		// Append struct to result slice
		result = append(result, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
