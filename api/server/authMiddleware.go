package server

import (
	"context"
	"net/http"
	"recipevault/api/util"
)

// Authenticate and authorize middleware
//
// Set the request context with important user details (userID)
func AuthMiddleware(log util.ILogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// Temp implementation
			// Place logic for decoding token here
			userID := 2

			ctx := context.WithValue(r.Context(), util.CTX_USER_ID, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
