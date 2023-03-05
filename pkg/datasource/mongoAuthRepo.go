package datasource

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AddAuthInfoArgs struct {
	ApiKey    string
	SecretKey string
}

func AddAuthInfo(client *mongo.Client, info AddAuthInfoArgs) {
	client.Database(MainDatabase).Collection(AuthInfoCollection).InsertOne(ctx, AuthInfo{
		SecretKey: info.SecretKey,
		ApiKey:    info.ApiKey,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
