package server

import (
	"net/http"
	"recipehub/api/src/util"
)

func AuthMiddleware(log util.ILogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
