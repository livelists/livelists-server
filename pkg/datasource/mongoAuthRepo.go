package datasource

import (
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddAuthInfoArgs struct {
	ApiKey    string
	SecretKey string
}

func AddAuthInfo(client *mongo.Client, info AddAuthInfoArgs) {
	client.Database(config.MainDatabase).Collection(mongoSchemes.AuthInfoCollection).InsertOne(
		ctx,
		mongoSchemes.NewAuthInfo(mongoSchemes.NewAuthInfoArgs{
			ApiKey:    info.ApiKey,
			SecretKey: info.SecretKey,
		}))
}
