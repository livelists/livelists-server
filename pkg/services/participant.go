package services

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/livelists/livelist-server/contracts/participant"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/services/accessToken"
)

type ParticipantService struct{}

func (p ParticipantService) AddParticipantToChannel(ctx context.Context, req *pb.AddParticipantToChannelReq) (*pb.AddParticipantToChannelRes, error) {
	part, err := datasource.AddParticipant(datasource.AddParticipantArgs{
		Identifier: req.Identifier,
		Channel:    req.ChannelId,
		Grants: pb.ChannelParticipantGrants{
			SendMessage:  req.Grants.SendMessage,
			Admin:        req.Grants.Admin,
			ReadMessages: req.Grants.ReadMessages,
		},
		CustomData: req.CustomData,
	})

	token := accessToken.AccessToken{}
	token.AddGrants(accessToken.GrantsData{
		SendMessage:  req.Grants.SendMessage,
		ReadMessages: req.Grants.ReadMessages,
		Admin:        req.Grants.Admin,
	})
	token.AddUser(req.Identifier)
	token.AddChannelId(req.ChannelId)

	tokenStr, err := token.Sign()

	fmt.Println(err)

	return &pb.AddParticipantToChannelRes{
		Participant: &pb.Participant{
			Identifier: part.Identifier,
			ChannelId:  part.Channel,
			CreatedAt:  &timestamp.Timestamp{Seconds: int64(part.CreatedAt.Second())},
			Status:     pb.ParticipantStatus_Active,
		},
		Grants: &pb.ChannelParticipantGrants{
			Admin:        &part.Grants.Admin,
			SendMessage:  &part.Grants.SendMessage,
			ReadMessages: &part.Grants.ReadMessages,
		},
		AccessToken: tokenStr,
	}, nil
}

func (p ParticipantService) GetParticipantAccessToken(ctx context.Context, req *pb.GetParticipantAccessTokenReq) (*pb.GetParticipantAccessTokenRes, error) {
	part, err := datasource.FindParticipantByIdentifierAndChannel(datasource.FindPByIdAndChannelArgs{
		Identifier: req.Identifier,
		ChannelId:  req.ChannelId,
	})

	token := accessToken.AccessToken{}

	token.AddGrants(accessToken.GrantsData{
		SendMessage:  &part.Grants.SendMessage,
		ReadMessages: &part.Grants.ReadMessages,
		Admin:        &part.Grants.Admin,
	})
	token.AddUser(part.Identifier)
	token.AddChannelId(req.ChannelId)

	tokenStr, err := token.Sign()

	if err != nil {
		fmt.Println(err, "token sign error")
	}

	return &pb.GetParticipantAccessTokenRes{
		Identifier:  req.Identifier,
		ChannelId:   req.ChannelId,
		AccessToken: tokenStr,
	}, nil
}
