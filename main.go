package main

import (
	"fmt"
	"github.com/ODawah/Distributed-URL-Shortener/Server"
	"github.com/ODawah/Distributed-URL-Shortener/models"
	"github.com/gin-gonic/gin"
	"net/http"
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

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ids)
}

func ShortenURL(ctx *gin.Context) {
	var url models.URL
	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println(url)
	ctx.JSON(http.StatusOK, url)
}
