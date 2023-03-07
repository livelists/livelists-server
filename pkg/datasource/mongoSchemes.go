package datasource

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const MainDatabase = "livelists"

const AuthInfoCollection = "AuthInfo"

type AuthInfo struct {
	ID        primitive.ObjectID `bson:"_id"`
	ApiKey    string             `bson:"apiKey"`
	SecretKey string             `bson:"secretKey"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

const ChannelCollection = "Channel"

type Channel struct {
	ID              primitive.ObjectID `bson:"_id"`
	Identification  string             `bson:"identification"`
	MaxParticipants int64              `bson:"maxParticipants"`
	Status          string             `bson:"status"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}
