package handler

import (
	"net/http"
	"recipevault/api/response"
	"recipevault/api/service"
	"recipevault/util"
)

// Get cookbook by id
func GetUserCookbook(log *util.Logger, s *service.CookbookService) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := GetCtxUserID(r)
		cookbookID := GetPathID(r, "cookbookID")

		data, err := s.GetUserCookbook(cookbookID, userID)

		// handle database error
		if err != nil {
			log.Error("handler->GetUserCookbook", err)
			response.Send500(w, "unable to retrieve cookbook")
			return
		}

		response.Send200(w, data)
	})
}
