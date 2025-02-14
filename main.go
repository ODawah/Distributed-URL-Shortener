package main

import (
	"github.com/ODawah/Distributed-URL-Shortener/Server"
)

func main() {
	r := Server.Routes()

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

}
