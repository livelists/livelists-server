package services

import (
	"github.com/livelists/livelist-server/pkg/datasource"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateChannel(client *mongo.Client) {
	datasource.CreateChannel(client)
}
