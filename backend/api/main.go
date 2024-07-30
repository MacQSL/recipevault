package main

import (
	"fmt"
	"net/http"
	"os"
	"recipehub/api/controllers"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/api", controllers.GetRoot)
	http.HandleFunc("/api/cookbooks", controllers.GetUserCookbooks)

	err := http.ListenAndServe(":9999", nil)

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
