package handler

import (
	"database/sql"
	"net/http"
	"recipevault/api/repository"
	"recipevault/api/response"
	"recipevault/api/service"
	"recipevault/util"
)

// Get all user cookbooks
func GetUserCookbooks(log *util.Logger, db *sql.DB) http.Handler {
	repository := repository.NewCookbookRepository(db)
	service := service.NewCookbookService(repository)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := GetCtxUserID(r)
		data, err := service.GetUserCookbooks(userID)

		if err != nil {
			response.Send500(w, err)
			return
		}

		response.Send200(w, data)
	})
}
