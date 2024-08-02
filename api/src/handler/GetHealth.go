package handler

import (
	"io"
	"net/http"
)

// Welcome to RecipeVault endpoint
func GetHealth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "RecipeVault Healthy 🍔")
}
