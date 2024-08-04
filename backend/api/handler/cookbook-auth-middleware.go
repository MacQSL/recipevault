package handler

import (
	"database/sql"
	"net/http"
	"recipevault/api/repository"
	"recipevault/api/service"
	"recipevault/util"
)

// Authorize user has access to cookbook
func CookbookAuthMiddleware(log util.ILogger, db *sql.DB) func(http.Handler) http.Handler {
	repository := repository.NewCookbookRepository(db)
	service := service.NewCookbookService(repository)

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			userID := getCtxUserID(r)
			cookbookID, err := getPathID(r, "cookbookID")

			if err != nil {
				log.Error("cookbook endpoint accessed without authority", err)
				w.WriteHeader(http.StatusForbidden)
				return
			}

			access := service.CanUserAccessCookbook(cookbookID, userID)

			if access {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "cookbook access denied", http.StatusForbidden)
			}

		}
		return http.HandlerFunc(fn)
	}
}
