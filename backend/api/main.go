package main

import (
	"fmt"
	"net/http"
	"os"
	"recipehub/api/controllers/get"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/api", get.Root)
	http.HandleFunc("/api/cookbooks", get.UserCookbooks)
	http.HandleFunc("/api/cookbooks/{id}", get.UserCookbooks)
	http.HandleFunc("/api/cookbooks/{id}/recipes", get.UserCookbooks)
	http.HandleFunc("/api/cookbooks/{id}/recipes/{id}", get.UserCookbooks)
	http.HandleFunc("/api/cookbooks/{id}/recipes/{id}/ingredients", get.UserCookbooks)
	http.HandleFunc("/api/cookbooks/{id}/recipes/{id}/ingredients{id}", get.UserCookbooks)

	err := http.ListenAndServe(":9999", nil)

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
