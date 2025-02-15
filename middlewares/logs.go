package middlewares

import (
	"context"
	"fmt"
	"github.com/ODawah/Distributed-URL-Shortener/models"
	"github.com/ODawah/Distributed-URL-Shortener/persistence"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LogRequestMiddleware(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if persistence.MongoClient == nil {
		log.Println("MongoDB client is not initialized")
		return
	}

	db := persistence.MongoClient.Database("client-requests")
	collection := db.Collection("requests")

	now := time.Now()
	formattedTime, err := time.Parse("02-Jan-2006 15:04:05", now.Format("02-Jan-2006 15:04:05"))
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}

	request := models.RequestData{
		ShortID:   ctx.Param("shortID"),
		Timestamp: formattedTime,
		IP:        ctx.ClientIP(),
	}

	// Insert the request data
	_, err = collection.InsertOne(c, request)
	if err != nil {
		log.Println("Failed to insert request data:", err)
	} else {
		fmt.Println("Request logged successfully:", request)
	}

	ctx.Next()
}
