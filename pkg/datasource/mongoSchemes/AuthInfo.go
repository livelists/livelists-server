package mongoSchemes

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const AuthInfoCollection = "AuthInfo"

type AuthInfo struct {
	ID        primitive.ObjectID `bson:"_id"`
	ApiKey    string             `bson:"apiKey"`
	SecretKey string             `bson:"secretKey"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type NewAuthInfoArgs struct {
	ApiKey    string
	SecretKey string
}

func NewAuthInfo(args NewAuthInfoArgs) AuthInfo {
	return AuthInfo{
		ID:        primitive.NewObjectID(),
		ApiKey:    args.ApiKey,
		SecretKey: args.SecretKey,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
