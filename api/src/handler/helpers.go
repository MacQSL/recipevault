package handler

import (
	"net/http"
	"recipevault/api/src/util"
	"strconv"
)

// Get userID from request context - see server/authMiddleware.go
func getCtxUserID(r *http.Request) int {
	id, ok := r.Context().Value(util.CTX_USER_ID).(int)

	// Invalid user ID
	if !ok {
		return -1
	}

	return id
}

// Parse ID from path by key
func parsePathID(r *http.Request, k string) (int, error) {
	return strconv.Atoi(r.PathValue(k))
}
