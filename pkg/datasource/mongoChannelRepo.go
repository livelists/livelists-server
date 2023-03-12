package datasource

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/livelists/livelist-server/contracts/channel"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"time"
)

type CreateChannelArgs struct {
	Identifier      string
	MaxParticipants int64
}

func CreateChannel(args CreateChannelArgs) pb.Channel {
	var client = config.GetMongoClient()

	newChannel := mongoSchemes.NewChannel(mongoSchemes.NewChannelArgs{
		Identifier:      args.Identifier,
		Status:          pb.ChannelStatus_Active,
		MaxParticipants: args.MaxParticipants,
	})

	_, err := client.Database(config.MainDatabase).Collection(mongoSchemes.ChannelCollection).InsertOne(ctx, newChannel)

	fmt.Print(err)

	return pb.Channel{
		Identifier:      args.Identifier,
		Status:          pb.ChannelStatus_Active,
		MaxParticipants: args.MaxParticipants,
		CreatedAt:       &timestamp.Timestamp{Seconds: int64(time.Now().Second())},
	}
}
