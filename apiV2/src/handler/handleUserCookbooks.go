package handler

import (
	"database/sql"
	"net/http"
	"recipevault/api/src/repository"
	"recipevault/api/src/service"
	"recipevault/api/src/util"
)

// Get all user cookbooks
func HandleUserCookbooks(log util.ILogger, db *sql.DB) http.HandlerFunc {
	repository := repository.NewCookbookRepository(log, db)
	service := service.NewCookbookService(log, repository)

	return func(w http.ResponseWriter, r *http.Request) {
		userID := getUserID(r)

		data, err := service.GetUserCookbooks(userID)

		if err != nil {
			log.Warn(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		err = encode(w, data)

		w.WriteHeader(http.StatusOK)
	}
}
