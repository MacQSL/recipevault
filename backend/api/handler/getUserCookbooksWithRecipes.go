package handler

import (
	"net/http"
	"recipevault/api/response"
	"recipevault/api/service"
	"recipevault/util"
)

// Get all user cookbooks
func GetUserCookbooksWithRecipes(log *util.Logger, s *service.CookbookService) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := GetCtxUserID(r)
		data, err := s.GetUserCookbooksWithRecipes(userID)

		if err != nil {
			log.Error("handler->GetUserCookbooks", err)
			response.Send500(w, "unable to retrieve cookbooks")
			return
		}

		response.Send200(w, data)
	})
}
