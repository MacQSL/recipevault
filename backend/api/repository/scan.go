package repository

import (
	"database/sql"
	"errors"
)

type Row[T any] interface {
	Ptrs() []any
	*T
}

func GetRow[T any, PT Row[T]](db *sql.DB, query string, args ...any) (T, error) {
	row := db.QueryRow(query, args...)

	var t T
	ptr := PT(&t)

	if err := row.Scan(ptr.Ptrs()...); err != nil {
		if err == sql.ErrNoRows {
			return t, errors.New("No row found")
		}
		return t, err
	}

	return t, nil
}

func FindRows[T any, PT Row[T]](db *sql.DB, query string, args ...any) ([]T, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var result []T

	for rows.Next() {
		var t T
		ptr := PT(&t)
		if err := rows.Scan(ptr.Ptrs()...); err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil

}
