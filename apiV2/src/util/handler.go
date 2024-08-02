package util

import "net/http"

// Collection of handler related utils

// Get userID from request context - see server/authMiddleware.go
func GetUserID(r *http.Request) int {
	return r.Context().Value("userID").(int)
}
