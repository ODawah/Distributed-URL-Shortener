package main

import (
	"errors"
	"fmt"
	"github.com/ODawah/Distributed-URL-Shortener/persistence"
	"github.com/ODawah/Distributed-URL-Shortener/server"
)

func main() {

	r := server.Routes()

	err := persistence.ConnectToRedis()
	if err != nil {
		panic(err)
	}

	_, _, err = persistence.ConnectToMongo()
	if err != nil {
		panic(err)
	}

	err = r.Run("0.0.0.0:8080")
	if err != nil {
		panic(errors.New(fmt.Sprintf("failed to run API: %v", err)))
	}

}
