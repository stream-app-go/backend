package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// VideoColl : Videos collection
var VideoColl *mongo.Collection

// CTX : db context
var CTX context.Context

// Connect : conntect to MongoDB instance
func Connect() {
	addr := os.Getenv("MONGO")
	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		panic(err)
	}
	CTX, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(CTX)
	if err != nil {
		panic(err)
	}
	VideoColl = client.Database("stream-app").Collection("videos")
}
