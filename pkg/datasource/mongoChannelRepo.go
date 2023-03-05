package datasource

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func CreateChannel(client *mongo.Client) {
	client.Database(MainDatabase).Collection(ChannelCollection).InsertOne(ctx, Channel{
		Identification: "nikita",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
}
