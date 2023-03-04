package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ctx = context.TODO()
var MongoClient *mongo.Client = nil

func ConnectToMongo(config MongoConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.URI)
	MongoClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = MongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return MongoClient, nil
}
