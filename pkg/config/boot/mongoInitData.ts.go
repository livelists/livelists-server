package boot

import (
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedMongo(client *mongo.Client, config *config.Config) {
	datasource.AddAuthInfo(client, datasource.AddAuthInfoArgs{
		ApiKey:    config.ApiKey,
		SecretKey: config.SecretKey,
	})
}
