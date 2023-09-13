package mongoSchemes

import (
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const MessageCollection = "Message"

type Message struct {
	ID         primitive.ObjectID `bson:"_id"`
	Channel    string             `bson:"channel"`
	Sender     primitive.ObjectID `bson:"participant"`
	Text       string             `bson:"text"`
	CustomData *map[string]string `bson:"customData"`
	Type       string             `bson:"type"`
	SubType    string             `bson:"subType"`
	DeletedAt  *time.Time         `bson:"deletedAt"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}

type NewMessageArgs struct {
	ChannelIdentifier string
	SenderId          primitive.ObjectID
	Text              string
	Type              string
	SubType           string
	CustomData        *wsMessages.CustomData
}

func NewMessage(args NewMessageArgs) Message {
	now := time.Now()
	if args.CustomData != nil {
		return Message{
			ID:         primitive.NewObjectID(),
			Channel:    args.ChannelIdentifier,
			Sender:     args.SenderId,
			Text:       args.Text,
			Type:       args.Type,
			SubType:    args.SubType,
			CustomData: &args.CustomData.Data,
			CreatedAt:  now,
			UpdatedAt:  now,
		}
	}
	return Message{
		ID:         primitive.NewObjectID(),
		Channel:    args.ChannelIdentifier,
		Sender:     args.SenderId,
		Text:       args.Text,
		Type:       args.Type,
		SubType:    args.SubType,
		CustomData: nil,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

type MessageWithParticipant struct {
	Id                string             `bson:"id"`
	ChannelIdentifier string             `bson:"channelIdentifier"`
	Text              string             `bson:"text"`
	CustomData        *map[string]string `bson:"customData"`
	Type              string             `bson:"type"`
	SubType           string             `bson:"subType"`
	CreatedAt         time.Time          `bson:"createdAt"`
	Participant       ShortParticipant   `bson:"participant"`
}
