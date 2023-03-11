package datasource

import (
	"context"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx = context.TODO()

var client = config.GetMongoClient()

type AddParticipantArgs struct {
	Identifier string
	Channel    primitive.ObjectID
	Grants     mongoSchemes.Grants
}

func AddParticipant(args AddParticipantArgs) (mongoSchemes.Participant, error) {
	newParticipant := mongoSchemes.NewParticipant(mongoSchemes.NewParticipantArgs{
		Identifier: args.Identifier,
		Channel:    args.Channel,
		Grants:     args.Grants,
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
