package main

import (
	"recipevault/api/src/server"

	_ "github.com/lib/pq"
)

func main() {
	server.Run()
}
