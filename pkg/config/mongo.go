package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ctx = context.TODO()

type MongoStore struct {
	DB *mongo.Client
}

var mStore MongoStore

func ConnectToMongo(config MongoConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.URI)
	mongoClient, err := mongo.Connect(ctx, clientOptions)

	mStore.DB = mongoClient

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return mongoClient, nil
}

func GetMongoClient() mongo.Client {
	if mStore.DB == nil {
		log.Printf("Mongo access before init")
		mStore.DB = &mongo.Client{}
	}
	return *mStore.DB
}
