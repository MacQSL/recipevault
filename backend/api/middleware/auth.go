package middleware

import (
	"context"
	"net/http"
	"recipevault/util"
)

// Authenticate and authorize middleware
//
// Set request context value for user ID
func AuthMiddleware(log *util.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Temp implementation, place logic for decoding token here
			userID := 2

			ctx := context.WithValue(r.Context(), util.CTX_USER_ID, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
