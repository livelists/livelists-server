package datasource

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/livelists/livelist-server/contracts/channel"
	"github.com/livelists/livelist-server/pkg/config"
	"time"
)

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

type CreateChannelArgs struct {
	Identification  string
	MaxParticipants int64
}

func CreateChannel(args CreateChannelArgs) pb.Channel {
	var client = config.GetMongoClient()
	_, err := client.Database(MainDatabase).Collection(ChannelCollection).InsertOne(ctx, Channel{
		Identification:  args.Identification,
		Status:          pb.ChannelStatus_Active.String(),
		MaxParticipants: args.MaxParticipants,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})

	fmt.Print(err)

	return pb.Channel{
		Identification:  args.Identification,
		Status:          0,
		MaxParticipants: args.MaxParticipants,
		CreatedAt:       &timestamp.Timestamp{Seconds: int64(time.Now().Second())},
	}
}
