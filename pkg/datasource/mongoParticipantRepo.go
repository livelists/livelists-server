package datasource

import (
	"context"
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

	newParticipant := mongoSchemes.NewParticipant(mongoSchemes.NewParticipantArgs{
		Identifier:        args.Identifier,
		ChannelIdentifier: args.Channel,
		Grants:            args.Grants,
		CustomData:        args.CustomData,
	})

	_, err := client.Database(config.MainDatabase).Collection(
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

	participant := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).FindOne(ctx, bson.D{
		{
			"identifier", args.Identifier,
		}, {
			"channel", args.ChannelId,
		}})

	var participantDocument mongoSchemes.Participant
	err := participant.Decode(&participantDocument)

	return participantDocument, err
}

type FindAllParticipantsByIdArgs struct {
	Identifier string
}

func FindAllParticipantsById(args FindAllParticipantsByIdArgs) ([]mongoSchemes.Participant, error) {
	var client = config.GetMongoClient()

	participant, _ := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).Find(ctx, bson.D{
		{
			"identifier", args.Identifier,
		}})

	var participantsDocuments []mongoSchemes.Participant
	err := participant.All(ctx, &participantsDocuments)

	return participantsDocuments, err
}

type UpdateParticipantLastSeenAtArgs struct {
	Identifier        string
	ChannelIdentifier string
	LastSeenAt        time.Time
	IsOnline          bool
}

func UpdateParticipantLastSeenAt(args UpdateParticipantLastSeenAtArgs) (time.Time, error) {
	var client = config.GetMongoClient()

	filter := bson.D{{
		"identifier", args.Identifier,
	}, {
		"channel", args.ChannelIdentifier,
	}}

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

	participants, err := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).Aggregate(ctx, bson.A{
		bson.D{{"$match", bson.D{{"channel", args.ChannelIdentifier}}}},
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

type UpdateLastMessageSeenAtArgs struct {
	LastSeenAtUnixMS  int64
	Identifier        string
	ChannelIdentifier string
}

func UpdateLastMessageSeenAt(args UpdateLastMessageSeenAtArgs) (time.Time, error) {
	var client = config.GetMongoClient()

	var createdAtTime = time.UnixMilli(args.LastSeenAtUnixMS)

	filter := bson.D{{
		"identifier", args.Identifier,
	}, {
		"channel", args.ChannelIdentifier,
	}}

	update := bson.D{{"$set", bson.D{{
		"lastSeenMessageCreatedAt", createdAtTime,
	}}}}

	var _, updateErr = client.Database(
		config.MainDatabase).Collection(mongoSchemes.ParticipantCollection).UpdateOne(
		ctx,
		filter,
		update)

	if updateErr != nil {
		return createdAtTime, updateErr
	}

	return createdAtTime, nil
}
