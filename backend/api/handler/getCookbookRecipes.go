package handler

import (
	"net/http"
	"recipevault/api/response"
	"recipevault/api/service"
	"recipevault/util"
)

// Get Cookbook Recipes
func GetCookbookRecipes(log *util.Logger, s *service.RecipeService) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookbookID := GetPathID(r, "cookbookID")

		data, err := s.GetCookbookRecipes(cookbookID)

		// handle database error
		if err != nil {
			log.Error("handler->GetCookbookRecipes", err)
			response.Send500(w, err, "unable to retrieve cookbook")
			return
		}

		response.Send200(w, data)
	})
}
