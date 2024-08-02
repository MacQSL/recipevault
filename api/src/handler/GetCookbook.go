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
func GetCookbook(log util.ILogger, db *sql.DB) http.HandlerFunc {
	repository := repository.NewCookbookRepository(db)
	service := service.NewCookbookService(repository)

	return func(w http.ResponseWriter, r *http.Request) {
		cookbookID, err := parsePathID(r, "cookbookID")

		// check path param is valid and integer
		if err != nil {
			log.Error("Invalid path param:", cookbookID)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := service.GetCookbook(cookbookID)

		// handle database error
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			log.Error("SQL error", err)
			w.WriteHeader(http.StatusInternalServerError)
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
