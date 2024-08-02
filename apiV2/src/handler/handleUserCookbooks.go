package handler

import (
	"database/sql"
	"net/http"
	"recipehub/api/src/repository"
	"recipehub/api/src/service"
	"recipehub/api/src/util"
)

// Get all user cookbooks
func HandleUserCookbooks(log util.ILogger, db *sql.DB) http.HandlerFunc {
	repository := repository.NewCookbookRepository(log, db)
	service := service.NewCookbookService(log, repository)

	return func(w http.ResponseWriter, r *http.Request) {
		userID := util.GetUserID(r)

		data, err := service.GetUserCookbooks(userID)

		if err != nil {
			log.Warn(err)
		}

		log.Info(data)
	}
}
