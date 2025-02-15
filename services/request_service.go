package services

import (
	"context"
	"fmt"
	"github.com/ODawah/Distributed-URL-Shortener/models"
	"github.com/ODawah/Distributed-URL-Shortener/persistence"
	"log"
	"time"
)

func LogRequestData(request models.RequestData) {
	// Set a timeout for MongoDB operations
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ensure MongoDB client is initialized
	if persistence.MongoClient == nil {
		log.Println("MongoDB client is not initialized")
		return
	}

	// Select the database and collection
	db := persistence.MongoClient.Database("client-requests")
	collection := db.Collection("requests")

	// Insert the request data
	_, err := collection.InsertOne(ctx, request)
	if err != nil {
		log.Println("Failed to insert request data:", err)
	} else {
		fmt.Println("Request logged successfully:", request)
	}
}
