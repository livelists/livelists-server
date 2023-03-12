package datasource

import (
	"context"
	"fmt"
	pb "github.com/livelists/livelist-server/contracts/participant"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.TODO()

type AddParticipantArgs struct {
	Identifier string
	Channel    primitive.ObjectID
	Grants     pb.ChannelParticipantGrants
}

func AddParticipant(args AddParticipantArgs) (mongoSchemes.Participant, error) {
	var client = config.GetMongoClient()

	channelIdBytes, _ := args.Channel.MarshalText()

	channel := client.Database(
		config.MainDatabase).Collection(mongoSchemes.ChannelCollection).FindOne(
		ctx, bson.D{{"identifier", string(channelIdBytes)}})

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
