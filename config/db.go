package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient  *mongo.Client
	DatabaseName = GetEnv("MONGODB_DBNAME")
)

// Max retries and delay settings
const (
	MaxRetries        = 5
	InitialRetryDelay = 2 * time.Second
)

func ConnectDatabase() {
	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configure MongoDB client options
	clientOptions := options.Client().ApplyURI(GetEnv("MONGODB_URI"))

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Verify connection with Ping
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	// Successful connection
	fmt.Println("Database Connected!")
	MongoClient = client
}

// DisconnectDatabase ensures proper cleanup
func DisconnectDatabase() {
	if MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := MongoClient.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from database: %v", err)
		} else {
			fmt.Println("Database Disconnected!")
		}
	}
}
