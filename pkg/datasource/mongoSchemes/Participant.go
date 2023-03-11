package mongoSchemes

import (
	pb "github.com/livelists/livelist-server/contracts/participant"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const ParticipantCollection = "Participant"

type Participant struct {
	ID         primitive.ObjectID `bson:"_id"`
	Identifier string             `bson:"identifier"`
	Status     string             `bson:"status"`
	Channel    primitive.ObjectID `bson:"channel"`
	Grants     Grants             `bson:"grants"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}

type Grants struct {
	SendMessage  *bool `bson:"sendMessage"`
	ReadMessages *bool `bson:"readMessages"`
	Admin        *bool `bson:"admin"`
}

type NewParticipantArgs struct {
	Identifier string
	Channel    primitive.ObjectID
	Grants     Grants
}

func NewParticipant(args NewParticipantArgs) Participant {
	return Participant{
		ID:         primitive.NewObjectID(),
		Identifier: args.Identifier,
		Channel:    args.Channel,
		Grants:     args.Grants,
		Status:     pb.ParticipantStatus_Active.String(),
		UpdatedAt:  time.Now(),
		CreatedAt:  time.Now(),
	}
}
