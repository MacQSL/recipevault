package handler

import (
	"io"
	"net/http"
)

// Get all cookbooks for a user
func GETuserCookbooks(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cookbooks")
}
