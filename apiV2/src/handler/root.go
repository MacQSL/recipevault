package handler

import (
	"io"
	"net/http"
)

// Welcome to RecipeHub endpoint
func Health(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "RecipeHub Healthy 🍔")
}
