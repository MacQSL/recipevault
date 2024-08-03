package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"recipevault/api/src/repository"
	"recipevault/api/src/service"
	"recipevault/api/src/util"
)

// Get all user cookbooks
func GetUserCookbooks(log util.ILogger, db *sql.DB) http.HandlerFunc {
	repository := repository.NewCookbookRepository(db)
	service := service.NewCookbookService(repository)

	return func(w http.ResponseWriter, r *http.Request) {
		userID := getCtxUserID(r)
		data, err := service.GetUserCookbooks(userID)

		if err != nil {
			log.Error("SQL error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jData, err := json.Marshal(data)

		if err != nil {
			log.Error("JSON parse", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jData)
	}
}
