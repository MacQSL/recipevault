package middleware

import (
	"database/sql"
	"net/http"
	"recipevault/api/handler"
	"recipevault/api/repository"
	"recipevault/api/response"
	"recipevault/api/service"
	"recipevault/util"
)

// Authorize user has access to cookbook
func CookbookAuthMiddleware(log *util.Logger, db *sql.DB) func(http.Handler) http.Handler {
	repository := repository.NewCookbookRepository(db)
	service := service.NewCookbookService(repository)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := handler.GetCtxUserID(r)
			cookbookID := handler.GetPathID(r, "cookbookID")

			// getPathID fallback value
			if cookbookID == -1 {
				response.Send400(w, "invalid cookbookID path param")
				return
			}

			access := service.CanUserAccessCookbook(cookbookID, userID)

			if access {
				next.ServeHTTP(w, r)
			} else {
				response.Send(w, http.StatusForbidden, "cookbook access denied")
			}
		})
	}
}
