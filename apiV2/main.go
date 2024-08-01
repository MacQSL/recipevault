package main

import (
	"recipehub/api/src/server"

	_ "github.com/lib/pq"
)

func main() {
	server.Run()
}
