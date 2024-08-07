package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMongodbClient() *mongo.Client {
	Client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017")) 
}
