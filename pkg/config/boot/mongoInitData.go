package boot

import (
	"context"
	confPackage "github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SeedMongo(client *mongo.Client, config *confPackage.Config) {
	db := client.Database(confPackage.MainDatabase)

	channelIndexModel := mongo.IndexModel{
		Keys:    bson.D{{"identifier", -1}},
		Options: options.Index().SetUnique(true),
	}
	db.Collection(mongoSchemes.ChannelCollection).Indexes().CreateOne(context.TODO(), channelIndexModel)

	messageChannelIndexModel := mongo.IndexModel{
		Keys: bson.D{{"channel", -1}},
	}
	db.Collection(mongoSchemes.MessageCollection).Indexes().CreateOne(context.TODO(), messageChannelIndexModel)

	messageParticipantIndexModel := mongo.IndexModel{
		Keys: bson.D{{"participant", -1}},
	}
	db.Collection(mongoSchemes.MessageCollection).Indexes().CreateOne(context.TODO(), messageParticipantIndexModel)

	datasource.AddAuthInfo(client, datasource.AddAuthInfoArgs{
		ApiKey:    config.ApiKey,
		SecretKey: config.SecretKey,
	})
}
