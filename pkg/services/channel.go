package services

import (
	"context"
	pb "github.com/livelists/livelist-server/contracts/channel"
	"github.com/livelists/livelist-server/pkg/datasource"
)

type ChannelService struct{}

func (s ChannelService) CreateChannel(ctx context.Context, req *pb.CreateChannelReq) (*pb.Channel, error) {
	ch := datasource.CreateChannel(datasource.CreateChannelArgs{
		Identifier:      req.Identifier,
		MaxParticipants: req.MaxParticipants,
	})

	return &pb.Channel{
		Identifier: ch.Identifier,
		CreatedAt:  ch.CreatedAt,
		Status:     ch.Status,
	}, nil
}
