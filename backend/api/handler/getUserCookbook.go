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
func GetUserCookbook(log util.ILogger, db *sql.DB) http.HandlerFunc {
	repository := repository.NewCookbookRepository(db)
	service := service.NewCookbookService(repository)

	return func(w http.ResponseWriter, r *http.Request) {
		userID := getCtxUserID(r)
		cookbookID, err := getPathID(r, "cookbookID")

		// check path param is valid and integer
		if err != nil {
			response.Send(w, http.StatusBadRequest, "invalid path param")
			return
		}

		data, err := service.GetUserCookbook(cookbookID, userID)

		// handle database error
		if err != nil {
			response.Send(w, http.StatusInternalServerError, "unable to retrieve cookbook")
			return
		}

		response.SendOk(w, data)
	}
}
