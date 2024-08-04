package main

import (
	"recipevault/api"

	_ "github.com/lib/pq"
)

func main() {
	api.StartServer()
}
