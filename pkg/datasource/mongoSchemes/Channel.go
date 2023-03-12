package mongoSchemes

import (
	"fmt"
	pb "github.com/livelists/livelist-server/contracts/channel"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const ChannelCollection = "Channel"

type Channel struct {
	ID              primitive.ObjectID `bson:"_id"`
	Identifier      string             `bson:"identifier"`
	MaxParticipants int64              `bson:"maxParticipants"`
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
	fmt.Println("channel args", args.Status.String())
	return Channel{
		ID:              primitive.NewObjectID(),
		Identifier:      args.Identifier,
		MaxParticipants: args.MaxParticipants,
		Status:          args.Status.String(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
