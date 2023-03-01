package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ctx = context.TODO()
var MongoClient *mongo.Client = nil

func ConnectToMongo(config MongoConfig) (*mongo.Client, error) {
	fmt.Println("try connect")
	clientOptions := options.Client().ApplyURI(config.URI)
	MongoClient, err := mongo.Connect(ctx, clientOptions)

	fmt.Println(err, "error")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = MongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("connected to mongo")

	return MongoClient, nil
}
