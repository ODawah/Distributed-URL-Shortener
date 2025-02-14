package main

import (
	"github.com/ODawah/Distributed-URL-Shortener/Server"
	"github.com/ODawah/Distributed-URL-Shortener/models"
)

var ids = []models.URL{
	{ID: "12345", URL: "www.google.com"},
	{ID: "12346", URL: "www.facebook.com"},
	{ID: "12347", URL: "www.twitter.com"},
}

func main() {
	r := Server.Routes()

	err := r.Run("localhost:8000")
	if err != nil {
		panic(err)
	}

}
