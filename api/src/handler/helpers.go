package handler

import (
	"net/http"
	"recipevault/api/src/util"
	"strconv"
)

// Get userID from request context - see server/authMiddleware.go
func getCtxUserID(r *http.Request) int {
	return r.Context().Value(util.CTX_USER_ID).(int)
}

// Parse ID from path by key
func parsePathID(r *http.Request, k string) (int, error) {
	return strconv.Atoi(r.PathValue(k))
}
