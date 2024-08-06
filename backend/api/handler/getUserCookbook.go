package handler

import (
	"database/sql"
	"net/http"
	"recipevault/api/repository"
	"recipevault/api/response"
	"recipevault/api/service"
	"recipevault/util"
)

// Get cookbook by id
func GetUserCookbook(log *util.Logger, db *sql.DB) http.Handler {
	repository := repository.NewCookbookRepository(db)
	service := service.NewCookbookService(repository)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := GetCtxUserID(r)
		cookbookID := GetPathID(r, "cookbookID")

		data, err := service.GetUserCookbook(cookbookID, userID)

		// handle database error
		if err != nil {
			response.Send(w, http.StatusInternalServerError, "unable to retrieve cookbook")
			return
		}

		response.Send200(w, data)
	})
}
