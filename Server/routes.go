package Server

import (
	"github.com/ODawah/Distributed-URL-Shortener/handlers"
	"github.com/ODawah/Distributed-URL-Shortener/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := GetRouter()

	r.Use(middlewares.Limiter)

	go middlewares.CleanupLimiters()

	r.GET("/:shortID", handlers.GetURL)

	r.POST("/shorten", handlers.ShortenURL)

	return r
}
