package persistence

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var MongoClient *mongo.Client

func ConnectToMongo() (database *mongo.Database, ctx context.Context, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("Failed to connect to Mongo: %v", err))

	}
	MongoClient = client

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("err", err)
		return
	}
	
	log.Println("Successfully connected to Mongo and pinged.")
	return
}
