package datasource

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/bson"
)

type AddMessageArgs struct {
	ChannelIdentifier string
	SenderIdentifier  string
	Text              string
	Type              string
	SubType           string
	CustomData        *wsMessages.CustomData
}

func AddMessage(args AddMessageArgs) (mongoSchemes.Message, error) {
	var client = config.GetMongoClient()

	channel := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ChannelCollection).FindOne(
		ctx, bson.D{{"identifier", args.ChannelIdentifier}})

	var channelDocument mongoSchemes.Channel
	err := channel.Decode(&channelDocument)

	if err != nil {
		fmt.Println("channel decode err")
		return mongoSchemes.Message{}, err
	}

	participant := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).FindOne(
		ctx, bson.D{{"identifier", args.SenderIdentifier}})

	var participantDocument mongoSchemes.Participant
	err = participant.Decode(&participantDocument)

	if err != nil {
		fmt.Println("participant decode err")
		return mongoSchemes.Message{}, err
	}

	newMessage := mongoSchemes.NewMessage(mongoSchemes.NewMessageArgs{
		ChannelId:  channelDocument.ID,
		SenderId:   participantDocument.ID,
		Text:       args.Text,
		Type:       args.Type,
		SubType:    args.SubType,
		CustomData: args.CustomData,
	})

	_, err = client.Database(
		config.MainDatabase).Collection(
		mongoSchemes.MessageCollection).InsertOne(ctx, newMessage)

	if err != nil {
		return mongoSchemes.Message{}, err
	}

	return newMessage, nil
}
