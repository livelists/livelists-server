package datasource

import (
	"context"
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
	client.Database("local").Collection("Channel").InsertOne(ctx, Channel{
		Identification: "nikita",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
}
