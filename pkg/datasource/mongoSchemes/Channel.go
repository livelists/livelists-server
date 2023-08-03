package mongoSchemes

import (
	pb "github.com/livelists/livelist-server/contracts/channel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const ChannelCollection = "Channel"

type Channel struct {
	ID              primitive.ObjectID `bson:"_id"`
	Identifier      string             `bson:"identifier"`
	MaxParticipants int64              `bson:"maxParticipants"`
	CustomData      *map[string]string `bson:"customData"`
	Status          string             `bson:"status"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}

type NewChannelArgs struct {
	Identifier      string
	MaxParticipants int64
	Status          pb.ChannelStatus
}

func NewChannel(args NewChannelArgs) Channel {
	return Channel{
		ID:              primitive.NewObjectID(),
		Identifier:      args.Identifier,
		MaxParticipants: args.MaxParticipants,
		Status:          args.Status.String(),
		CustomData:      nil,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

type ChannelWithLastMessages struct {
	Channel  shortChannel             `bson:"channel"`
	Messages []MessageWithParticipant `bson:"messages"`
}

type shortChannel struct {
	Id         string             `bson:"id"`
	Identifier string             `bson:"identifier"`
	Status     string             `bson:"status"`
	CustomData *map[string]string `bson:"customData"`
}
