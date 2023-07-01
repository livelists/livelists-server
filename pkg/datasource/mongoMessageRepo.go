package datasource

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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

	var channelDocument, err = FindChannelByIdentifier(args.ChannelIdentifier)

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

func FindChannelByIdentifier(identifier string) (mongoSchemes.Channel, error) {
	var client = config.GetMongoClient()

	channel := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ChannelCollection).FindOne(
		ctx, bson.D{{"identifier", identifier}})

	var channelDocument mongoSchemes.Channel
	err := channel.Decode(&channelDocument)

	if err != nil {
		return mongoSchemes.Channel{}, err
	}

	return channelDocument, nil
}

type GetMessagesFromChannelArgs struct {
	ChannelIdentifier string
	Skip              int
	Limit             int
	StartFromDate     time.Time
}

func GetMessagesFromChannel(args GetMessagesFromChannelArgs) ([]mongoSchemes.MessageWithParticipant, int64, error) {
	var client = config.GetMongoClient()

	var channelDocument, err = FindChannelByIdentifier(args.ChannelIdentifier)

	if err != nil {
		return []mongoSchemes.MessageWithParticipant{}, 0, err
	}

	messages, err := client.Database(
		config.MainDatabase).Collection(mongoSchemes.MessageCollection).Aggregate(ctx, bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"channel", channelDocument.ID},
					{"createdAt", bson.D{{"$lt", args.StartFromDate}}},
				},
			},
		},
		bson.D{{"$sort", bson.D{{"createdAt", -1}}}},
		bson.D{{"$skip", args.Skip}},
		bson.D{{"$limit", args.Limit}},
		bson.D{{"$sort", bson.D{{"createdAt", 1}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "Participant"},
					{"localField", "participant"},
					{"foreignField", "_id"},
					{"as", "participant"},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"id", "$_id"},
					{"participant",
						bson.D{
							{"$arrayElemAt",
								bson.A{
									"$participant",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"_id", 0},
					{"id", 1},
					{"text", 1},
					{"customData", 1},
					{"type", 1},
					{"subType", 1},
					{"participant.identifier", 1},
					{"participant.customData", 1},
					{"createdAt", 1},
				},
			},
		},
	})

	var messagesDocuments []mongoSchemes.MessageWithParticipant

	err = messages.All(ctx, &messagesDocuments)

	if err != nil {
		return []mongoSchemes.MessageWithParticipant{}, 0, err
	}

	var totalCount, errCount = CountMessagesInChannel(CountMessagesInChannelArgs{
		ChannelId: channelDocument.ID,
	})

	if errCount != nil {
		return []mongoSchemes.MessageWithParticipant{}, 0, errCount
	}

	return messagesDocuments, totalCount, nil
}

type CountMessagesInChannelArgs struct {
	ChannelId primitive.ObjectID
}

func CountMessagesInChannel(args CountMessagesInChannelArgs) (int64, error) {
	var client = config.GetMongoClient()

	messagesCount, err := client.Database(config.MainDatabase).Collection(mongoSchemes.MessageCollection).CountDocuments(ctx, bson.D{{"channel", args.ChannelId}})

	if err != nil {
		fmt.Println("Count messages error")
		return 0, err
	}

	fmt.Println("Count messages", messagesCount)

	return messagesCount, nil
}
