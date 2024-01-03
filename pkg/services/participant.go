package services

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/livelists/livelist-server/contracts/participant"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/services/accessToken"
	"github.com/livelists/livelist-server/pkg/services/message"
	"github.com/livelists/livelist-server/pkg/shared"
)

type ParticipantService struct {
	WS shared.WsRoom
}

func (p ParticipantService) AddParticipantToChannel(ctx context.Context, req *pb.AddParticipantToChannelReq) (*pb.AddParticipantToChannelRes, error) {
	if ctx.Value("isTokenValid").(bool) == false {
		return &pb.AddParticipantToChannelRes{
			Errors: []pb.AddParticipantToChannelErrors{pb.AddParticipantToChannelErrors_Unauthorized},
		}, nil
	}

	existedParticipant := datasource.FindParticipantByIdentifierAndChannel(datasource.FindPByIdAndChannelArgs{
		Identifier: req.Identifier,
		ChannelId:  req.ChannelId,
	})

	if existedParticipant != nil {
		return &pb.AddParticipantToChannelRes{
			Errors: []pb.AddParticipantToChannelErrors{pb.AddParticipantToChannelErrors_IsAlreadyExist},
		}, nil
	}

	if len(req.Identifier) == 0 {
		return &pb.AddParticipantToChannelRes{
			Errors: []pb.AddParticipantToChannelErrors{pb.AddParticipantToChannelErrors_IdentifierNotValid},
		}, nil
	}

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

	tokenStr, err := token.Sign()

	if err == nil {
		message.CreateMessage(&message.CreateMessageArgs{
			Payload: message.CreateMessagePayload{
				Text:    req.Identifier,
				LocalId: "-",
			},
			ChannelId:        req.ChannelId,
			SenderIdentifier: &part.Identifier,
			WS:               p.WS,
			Type:             wsMessages.MessageType_System,
			SubType:          wsMessages.MessageSubType_ParticipantJoined,
		})
	}

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
	if ctx.Value("isTokenValid").(bool) == false {
		return &pb.GetParticipantAccessTokenRes{
			Errors: []pb.GetParticipantAccessTokenErrors{pb.GetParticipantAccessTokenErrors_RootUnauthorized},
		}, nil
	}

	part := datasource.FindParticipantByIdentifierAndChannel(datasource.FindPByIdAndChannelArgs{
		Identifier: req.Identifier,
		ChannelId:  req.ChannelId,
	})

	if part == nil {
		return &pb.GetParticipantAccessTokenRes{
			Errors: []pb.GetParticipantAccessTokenErrors{pb.GetParticipantAccessTokenErrors_NotFound},
		}, nil
	}

	token := accessToken.AccessToken{}

	token.AddGrants(accessToken.GrantsData{
		SendMessage:  &part.Grants.SendMessage,
		ReadMessages: &part.Grants.ReadMessages,
		Admin:        &part.Grants.Admin,
	})
	token.AddUser(part.Identifier)

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
