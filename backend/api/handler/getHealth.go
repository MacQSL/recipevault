package handler

import (
	"io"
	"net/http"
)

// Welcome to RecipeVault endpoint
func GetHealth(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "RecipeVault Healthy 🍔")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
