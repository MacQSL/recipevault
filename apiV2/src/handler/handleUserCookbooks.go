package handler

import (
	"database/sql"
	"net/http"
	"recipehub/api/src/repository"
	"recipehub/api/src/service"
	"recipehub/api/src/util"
)

func HandleUserCookbooks(log util.ILogger, db *sql.DB) http.HandlerFunc {
	repository := repository.NewCookbookRepository(db)
	service := service.NewCookbookService(repository)

	return func(w http.ResponseWriter, r *http.Request) {
		data, err := service.GetUserCookbooks(2)

		if err != nil {
			log.Warn(err)
		}

		log.Info(data)
	}
}
