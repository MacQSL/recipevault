package handler

import (
	"net/http"
	"recipevault/api/response"
)

// RecipeVault health check
func GetHealth() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response.Send200(w, "RecipeVault healthy")
	})
}
