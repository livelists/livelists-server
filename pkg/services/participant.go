package services

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/livelists/livelist-server/contracts/participant"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ParticipantService struct{}

func (p ParticipantService) AddParticipantToChannel(ctx context.Context, req *pb.AddParticipantToChannelReq) (*pb.ParticipantCreateRes, error) {
	channelObjId, _ := primitive.ObjectIDFromHex(req.ChannelId)

	part, err := datasource.AddParticipant(datasource.AddParticipantArgs{
		Identifier: req.Identifier,
		Channel:    channelObjId,
		Grants: mongoSchemes.Grants{
			SendMessage:  req.Grants.SendMessage,
			Admin:        req.Grants.Admin,
			ReadMessages: req.Grants.ReadMessages,
		},
	})

	fmt.Println(err)

	return &pb.ParticipantCreateRes{
		Participant: &pb.Participant{
			Identifier: part.Identifier,
			ChannelId:  part.Channel.String(),
			CreatedAt:  &timestamp.Timestamp{Seconds: int64(part.CreatedAt.Second())},
			Status:     pb.ParticipantStatus_Active,
		},
		Grants: &pb.ChannelParticipantGrants{
			Admin:        part.Grants.Admin,
			SendMessage:  part.Grants.SendMessage,
			ReadMessages: part.Grants.ReadMessages,
		},
		AccessToken: "fgfh",
	}, nil
}
