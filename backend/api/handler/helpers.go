package handler

import (
	"net/http"
	"recipevault/util"
	"strconv"
)

// Get userID from request context [server.authMiddleware]
func GetCtxUserID(r *http.Request) int {
	id, ok := r.Context().Value(util.CTX_USER_ID).(int)
	if !ok {
		return -1 // Invalid user ID
	}
	return id
}

// Parse ID from URL path by key
func GetPathID(r *http.Request, k string) int {
	id, err := strconv.Atoi(r.PathValue(k))
	if err != nil {
		return -1 // Invalid ID
	}
	return id
}
