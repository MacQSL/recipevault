package handler

import (
	"io"
	"net/http"
)

// Get all cookbooks for a user
func UserCookbooks(res http.ResponseWriter, req *http.Request) {
	//var connection = database.Connect()
	io.WriteString(res, "cookbooks")
}
