package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recipevault/api/src/util"
)

// Collection of handler related utils

// Get userID from request context - see server/authMiddleware.go
func getUserID(r *http.Request) int {
	return r.Context().Value(util.CTX_USER_ID).(int)
}

// Encode the response as JSON
func encode[T any](w http.ResponseWriter, v T) error {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

// Decode the request body as JSON
func decode[T any](r *http.Request) (T, error) {
	var v T

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	return v, nil
}
