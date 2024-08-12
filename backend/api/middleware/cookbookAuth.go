package middleware

import (
	"errors"
	"net/http"
	"recipevault/api/handler"
	"recipevault/api/response"
	"recipevault/api/service"
	"recipevault/util"
)

// Authorize user has access to cookbook
func CookbookAuthMiddleware(log *util.Logger, s *service.CookbookService) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := handler.GetCtxUserID(r)
			cookbookID := handler.GetPathID(r, "cookbookID")

			if cookbookID == -1 {
				response.Send400(w, errors.New("cannot retrieve path param"), "invalid cookbookID path param")
				return
			}

			// NOTE: There might be a optimization where we include the
			// user cookbook access within the token. I think the performance
			// gains would be very trivial, but would prevent having to call the database
			// on every cookbook ID related endpoint before the actual query
			access := s.CanUserAccessCookbook(cookbookID, userID)

			if access {
				next.ServeHTTP(w, r)
			} else {
				response.Send403(w, errors.New("cookbook auth failed"), "cookbook access denied")
			}
		})
	}
}
