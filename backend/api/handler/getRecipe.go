package handler

import (
	"net/http"
	"recipevault/api/response"
	"recipevault/api/service"
	"recipevault/util"
)

// Get recipe by id
func GetRecipe(log *util.Logger, s *service.RecipeService) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recipeID := GetPathID(r, "recipeID")

		data, err := s.GetRecipe(recipeID)

		// handle database error
		if err != nil {
			log.Error("handler->GetRecipe:", err)
			response.SendError(w, err, "unable to retrieve recipe")
			return
		}

		response.Send200(w, data)
	})
}
