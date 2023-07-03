package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartClient() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.
		Client().
		ApplyURI("TODO").
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}

	// Ping database to confirm connection.
	err = client.
		Database("LoggingMiddleware").
		RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).
		Err()
	if err != nil {
		return err
	}

	return nil
}
