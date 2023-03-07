package services

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/livelists/livelist-server/contracts/channel"
	"github.com/livelists/livelist-server/pkg/datasource"
)

type ChannelService struct {
}

func (s ChannelService) CreateChannel(ctx context.Context, req *pb.CreateChannelReq) (*pb.Channel, error) {
	datasource.CreateChannel(datasource.CreateChannelArgs{
		Identification:  req.Identification,
		MaxParticipants: req.MaxParticipants,
	})

	return &pb.Channel{Identification: "Channel", CreatedAt: &timestamp.Timestamp{Seconds: 0}, Status: pb.ChannelStatus_Active}, nil
}
