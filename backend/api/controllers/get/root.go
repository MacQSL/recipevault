package get

import (
	"io"
	"net/http"
)

// Welcome to RecipeHub endpoint
func Root(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to RecipeHub API.")
}
