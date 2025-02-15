package main

import (
	"github.com/ODawah/Distributed-URL-Shortener/Server"
	"github.com/ODawah/Distributed-URL-Shortener/persistence"
)

func main() {

	r := Server.Routes()

	err := persistence.ConnectToRedis()
	if err != nil {
		panic(err)
	}

	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

}
