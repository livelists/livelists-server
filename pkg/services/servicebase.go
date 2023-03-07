package services

import "go.mongodb.org/mongo-driver/mongo"

type ServiceBase struct {
	mongo *mongo.Client
}
