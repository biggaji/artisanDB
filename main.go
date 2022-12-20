package main

import (
	"api"
	"db"
)

func main() {
	dbase := db.MakeDB()

	server := api.MakeDBServer(dbase)

	server.Start()
}
