package datasource

import (
	"context"
	"fmt"
	pb "github.com/livelists/livelist-server/contracts/participant"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

var ctx = context.TODO()

type AddParticipantArgs struct {
	Identifier string
	Channel    string
	Grants     pb.ChannelParticipantGrants
	CustomData *pb.CustomData
}

func AddParticipant(args AddParticipantArgs) (mongoSchemes.Participant, error) {
	var client = config.GetMongoClient()

	channel := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ChannelCollection).FindOne(
		ctx, bson.D{{"identifier", args.Channel}})

	var channelDocument mongoSchemes.Channel
	err := channel.Decode(&channelDocument)

	if err != nil {
		fmt.Println("channel decode err")
		return mongoSchemes.Participant{}, err
	}

	newParticipant := mongoSchemes.NewParticipant(mongoSchemes.NewParticipantArgs{
		Identifier: args.Identifier,
		ChannelId:  channelDocument.ID,
		Grants:     args.Grants,
		CustomData: args.CustomData,
	})

	_, err = client.Database(config.MainDatabase).Collection(
		mongoSchemes.ParticipantCollection).InsertOne(ctx, newParticipant)

	return mongoSchemes.Participant{
		Identifier: newParticipant.Identifier,
		Channel:    newParticipant.Channel,
		Grants:     newParticipant.Grants,
		CreatedAt:  newParticipant.CreatedAt,
		Status:     newParticipant.Status,
	}, err
}

type FindPByIdAndChannelArgs struct {
	Identifier string
	ChannelId  string
}

func FindParticipantByIdentifierAndChannel(args FindPByIdAndChannelArgs) (mongoSchemes.Participant, error) {
	var client = config.GetMongoClient()

	channel := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ChannelCollection).FindOne(
		ctx, bson.D{{"identifier", args.ChannelId}})

	var channelDocument mongoSchemes.Channel
	err := channel.Decode(&channelDocument)

	if err != nil {
		return mongoSchemes.Participant{}, err
	}

	participant := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).FindOne(ctx, bson.D{
		{
			"identifier", args.Identifier,
		}, {
			"channel", channelDocument.ID,
		}})

	var participantDocument mongoSchemes.Participant
	err = participant.Decode(&participantDocument)

	return participantDocument, err
}

type UpdateParticipantLastSeenAtArgs struct {
	Identifier        string
	ChannelIdentifier string
	LastSeenAt        time.Time
	IsOnline          bool
}

func UpdateParticipantLastSeenAt(args UpdateParticipantLastSeenAtArgs) (time.Time, error) {
	var client = config.GetMongoClient()

	var participant, err = FindParticipantByIdentifierAndChannel(FindPByIdAndChannelArgs{
		ChannelId:  args.ChannelIdentifier,
		Identifier: args.Identifier,
	})

	if err != nil {
		return args.LastSeenAt, err
	}

	filter := bson.D{{"_id", participant.ID}}
	update := bson.D{{"$set", bson.D{{
		"lastSeenAt", args.LastSeenAt,
	}, {"isOnline", args.IsOnline}}}}

	var _, updateErr = client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).UpdateOne(
		ctx,
		filter,
		update)

	if updateErr != nil {
		return args.LastSeenAt, updateErr
	}

	return args.LastSeenAt, nil
}

type GetShortParticipantsArgs struct {
	ChannelIdentifier string
	Limit             int32
}

func GetShortParticipants(args GetShortParticipantsArgs) ([]mongoSchemes.ShortParticipant, error) {
	var client = config.GetMongoClient()

	var channelDocument, err = FindChannelByIdentifier(args.ChannelIdentifier)

	if err != nil {
		return []mongoSchemes.ShortParticipant{}, err
	}

	participants, err := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).Aggregate(ctx, bson.A{
		bson.D{{"$match", bson.D{{"channel", channelDocument.ID}}}},
		bson.D{
			{"$sort",
				bson.D{
					{"isOnline", 1},
					{"lastSeenAt", 1},
				},
			},
		},
		bson.D{{"$limit", args.Limit}},
	})

	var participantsDocuments []mongoSchemes.ShortParticipant

	err = participants.All(ctx, &participantsDocuments)

	if err != nil {
		return []mongoSchemes.ShortParticipant{}, err
	}

	return participantsDocuments, nil
}
