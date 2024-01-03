package mongoSchemes

import (
	pb "github.com/livelists/livelist-server/contracts/channel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const ChannelCollection = "Channel"

type Channel struct {
	ID                       primitive.ObjectID `bson:"_id"`
	Identifier               string             `bson:"identifier"`
	MaxParticipants          int64              `bson:"maxParticipants"`
	CustomData               *map[string]string `bson:"customData"`
	Status                   string             `bson:"status"`
	LastSeenMessageCreatedAt time.Time          `bson:"lastSeenMessageCreatedAt"`
	CreatedAt                time.Time          `bson:"createdAt"`
	UpdatedAt                time.Time          `bson:"updatedAt"`
}

type NewChannelArgs struct {
	Identifier      string
	MaxParticipants int64
	Status          pb.ChannelStatus
	CustomData      *pb.CustomData
}

func NewChannel(args NewChannelArgs) Channel {
	var customData map[string]string

	if args.CustomData != nil {
		customData = args.CustomData.Data
	}

	return Channel{
		ID:                       primitive.NewObjectID(),
		Identifier:               args.Identifier,
		MaxParticipants:          args.MaxParticipants,
		Status:                   args.Status.String(),
		CustomData:               &customData,
		LastSeenMessageCreatedAt: time.Now(),
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}
}

type ChannelWithLastMessages struct {
	Channel     shortChannel             `bson:"channel"`
	UnreadCount int64                    `bson:"unreadCount"`
	Messages    []MessageWithParticipant `bson:"messages"`
}

type shortChannel struct {
	Id         string             `bson:"id"`
	Identifier string             `bson:"identifier"`
	Status     string             `bson:"status"`
	CreatedAt  time.Time          `bson:"createdAt"`
	CustomData *map[string]string `bson:"customData"`
}

type ChannelParticipantsCount struct {
	ParticipantsCount       int64   `bson:"participantsCount"`
	OnlineParticipantsCount int64   `bson:"onlineParticipantsCount"`
	Channel                 Channel `bson:"channel"`
}
