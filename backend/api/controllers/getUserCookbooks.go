package controllers

import (
	"io"
	"net/http"
)

func GetUserCookbooks(res http.ResponseWriter, req *http.Request) {
	//var connection = database.Connect()

	io.WriteString(res, "Welcome to RecipeHub API.")
}
