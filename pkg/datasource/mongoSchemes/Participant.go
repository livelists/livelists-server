package mongoSchemes

import (
	"fmt"
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
	CustomData *map[string]string `bson:"customData"`
	Grants     Grants             `bson:"grants"`
	CreatedAt  time.Time          `bson:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}

type Grants struct {
	SendMessage  bool `bson:"sendMessage"`
	ReadMessages bool `bson:"readMessages"`
	Admin        bool `bson:"admin"`
}

type NewParticipantArgs struct {
	Identifier string
	ChannelId  primitive.ObjectID
	Grants     pb.ChannelParticipantGrants
	CustomData *pb.CustomData
}

func NewParticipant(args NewParticipantArgs) Participant {
	now := time.Now()

	fmt.Println("custom1", args.CustomData)
	if args.CustomData != nil {
		fmt.Println("custom Data", &args.CustomData.Data)
		return Participant{
			ID:         primitive.NewObjectID(),
			Identifier: args.Identifier,
			Channel:    args.ChannelId,
			CustomData: &args.CustomData.Data,
			Grants: Grants{
				Admin:        FalseIfNil(args.Grants.Admin),
				SendMessage:  FalseIfNil(args.Grants.SendMessage),
				ReadMessages: FalseIfNil(args.Grants.ReadMessages),
			},
			Status:    pb.ParticipantStatus_Active.String(),
			UpdatedAt: now,
			CreatedAt: now,
		}
	}
	return Participant{
		ID:         primitive.NewObjectID(),
		Identifier: args.Identifier,
		Channel:    args.ChannelId,
		CustomData: nil,
		Grants: Grants{
			Admin:        FalseIfNil(args.Grants.Admin),
			SendMessage:  FalseIfNil(args.Grants.SendMessage),
			ReadMessages: FalseIfNil(args.Grants.ReadMessages),
		},
		Status:    pb.ParticipantStatus_Active.String(),
		UpdatedAt: now,
		CreatedAt: now,
	}
}
