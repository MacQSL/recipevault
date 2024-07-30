package controllers

import (
	"io"
	"net/http"
)

func GetRoot(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome to RecipeHub API.")
}
