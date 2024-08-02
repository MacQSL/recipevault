package handler

import (
	"database/sql"
	"net/http"
	"recipehub/api/src/repository"
	"recipehub/api/src/service"
	"recipehub/api/src/util"
)

func HandleUserCookbooks(log *util.Logger, db *sql.DB) http.HandlerFunc {
	cr := repository.NewCookbookRepository(db)
	cs := service.NewCookbookService(cr)

	return func(w http.ResponseWriter, r *http.Request) {
		data, err := cs.GetUserCookbooks(2)

		if err != nil {
			log.Warn(err)
		}

		log.Info(data)
	}
}
