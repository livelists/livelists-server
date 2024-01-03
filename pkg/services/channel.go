package services

import (
	"context"
	pb "github.com/livelists/livelist-server/contracts/channel"
	"github.com/livelists/livelist-server/pkg/datasource"
)

type ChannelService struct{}

func (s ChannelService) CreateChannel(ctx context.Context, req *pb.CreateChannelReq) (*pb.CreateChannelResponse, error) {
	if ctx.Value("isTokenValid").(bool) == false {
		return &pb.CreateChannelResponse{
			Errors: []pb.CreateChannelErrors{pb.CreateChannelErrors_Unauthorized},
		}, nil
	}

	var channel = datasource.FindChannelByIdentifier(datasource.FindChannelByIdentifierArgs{
		Identifier: req.Identifier,
	})

	if *&channel != nil {
		return &pb.CreateChannelResponse{
			Errors: []pb.CreateChannelErrors{pb.CreateChannelErrors_IsAlreadyExist},
		}, nil
	}

	ch := datasource.CreateChannel(datasource.CreateChannelArgs{
		Identifier:      req.Identifier,
		MaxParticipants: req.MaxParticipants,
		CustomData:      req.CustomData,
	})

	return &pb.CreateChannelResponse{
		Channel: &pb.Channel{
			Identifier:      ch.Identifier,
			CreatedAt:       ch.CreatedAt,
			Status:          ch.Status,
			MaxParticipants: ch.MaxParticipants,
			CustomData:      ch.GetCustomData(),
		},
	}, nil
}
