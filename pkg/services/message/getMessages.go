package message

import (
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"time"
)

type GetMessagesArgs struct {
	PageSize          int
	Offset            int
	ChannelIdentifier string
	StartFromDate     time.Time
}

func GetMessages(args GetMessagesArgs) ([]mongoSchemes.MessageWithParticipant, error) {
	messages, err := datasource.GetMessagesFromChannel(datasource.GetMessagesFromChannelArgs{
		ChannelIdentifier: args.ChannelIdentifier,
		Limit:             args.PageSize,
		Skip:              args.Offset,
		StartFromDate:     args.StartFromDate,
	})
	return messages, err
}
