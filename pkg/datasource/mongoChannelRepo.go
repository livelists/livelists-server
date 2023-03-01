package datasource

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

type Channel struct {
	ID             primitive.ObjectID `bson:"_id"`
	Identification string             `bson:"identification""`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}

func CreateChannel(client *mongo.Client) {
	fmt.Println("create channel")
	databases, _ := client.ListDatabaseNames(ctx, ctx)

	fmt.Println(databases)

	client.Database("local").Collection("Channel").InsertOne(ctx, Channel{
		Identification: "nikita",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
}
