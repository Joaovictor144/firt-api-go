package main

import (
	"first-api-go/db"
	server2 "first-api-go/server"
)

func main() {
	db.StartDB()

	server  := server2.NewServer()

	server.Run()
}
