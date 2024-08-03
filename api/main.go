package main

import (
	"recipevault/api/server"

	_ "github.com/lib/pq"
)

func main() {
	server.Run()
}
