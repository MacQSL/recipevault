package handler

import (
	"io"
	"net/http"
)

// Welcome to RecipeHub endpoint
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "RecipeHub Healthy 🍔")
}
