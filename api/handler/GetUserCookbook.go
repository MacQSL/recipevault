package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"recipevault/api/src/repository"
	"recipevault/api/src/service"
	"recipevault/api/src/util"
)

// Get cookbook by id
func GetUserCookbook(log util.ILogger, db *sql.DB) http.HandlerFunc {
	repository := repository.NewCookbookRepository(db)
	service := service.NewCookbookService(repository)

	return func(w http.ResponseWriter, r *http.Request) {
		userID := getCtxUserID(r)
		cookbookID, err := parsePathID(r, "cookbookID")

		// check path param is valid and integer
		if err != nil {
			http.Error(w, "invalid path param", http.StatusBadRequest)
			return
		}

		data, err := service.GetUserCookbook(cookbookID, userID)

		// handle database error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jData, err := json.Marshal(data)

		// handle json parsing errors
		if err != nil {
			log.Error("JSON parse", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jData)
	}
}
