package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.
		Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	// Ping database to confirm connection.
	err = client.
		Database("LoggingMiddleware").
		RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).
		Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
