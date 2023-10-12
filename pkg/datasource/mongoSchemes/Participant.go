package mongoSchemes

import (
	pb "github.com/livelists/livelist-server/contracts/participant"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const ParticipantCollection = "Participant"

type ShortParticipant struct {
	ID         string             `bson:"id"`
	Identifier string             `bson:"identifier"`
	CustomData *map[string]string `bson:"customData"`
	LastSeenAt time.Time          `bson:"lastSeenAt"`
	IsOnline   bool               `bson:"isOnline"`
}

type Participant struct {
	ID                       primitive.ObjectID `bson:"_id"`
	Identifier               string             `bson:"identifier"`
	Status                   string             `bson:"status"`
	Channel                  string             `bson:"channel"`
	CustomData               *map[string]string `bson:"customData"`
	Grants                   Grants             `bson:"grants"`
	LastSeenAt               time.Time          `bson:"lastSeenAt"`
	LastSeenMessageCreatedAt time.Time          `bson:"lastSeenMessageCreatedAt"`
	IsOnline                 bool               `bson:"isOnline"`
	CreatedAt                time.Time          `bson:"createdAt"`
	UpdatedAt                time.Time          `bson:"updatedAt"`
}

type Grants struct {
	SendMessage  bool `bson:"sendMessage"`
	ReadMessages bool `bson:"readMessages"`
	Admin        bool `bson:"admin"`
}

type NewParticipantArgs struct {
	Identifier        string
	ChannelIdentifier string
	Grants            pb.ChannelParticipantGrants
	CustomData        *pb.CustomData
}

func NewParticipant(args NewParticipantArgs) Participant {
	now := time.Now()

	if args.CustomData != nil {
		return Participant{
			ID:                       primitive.NewObjectID(),
			Identifier:               args.Identifier,
			Channel:                  args.ChannelIdentifier,
			CustomData:               &args.CustomData.Data,
			LastSeenAt:               time.Unix(0, 0),
			LastSeenMessageCreatedAt: time.Now(),
			IsOnline:                 false,
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
		Channel:    args.ChannelIdentifier,
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
