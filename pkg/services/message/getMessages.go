package message

import (
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/shared/helpers"
	"time"
)

type GetMessagesArgs struct {
	PageSize          int
	Offset            int
	ChannelIdentifier string
	StartFromDate     time.Time
}

type GetMessagesRes struct {
	Messages              []*wsMessages.Message
	TotalCount            int64
	FirstMessageCreatedAt time.Time
	LastMessageCreatedAt  time.Time
}

func GetMessages(args GetMessagesArgs) (GetMessagesRes, error) {
	result, err := datasource.GetMessagesFromChannel(datasource.GetMessagesFromChannelArgs{
		ChannelIdentifier: args.ChannelIdentifier,
		Limit:             args.PageSize,
		Skip:              args.Offset,
		StartFromDate:     args.StartFromDate,
	})

	return GetMessagesRes{
		Messages:              helpers.MongoMessagesToPB(*result.Messages),
		TotalCount:            result.TotalCount,
		FirstMessageCreatedAt: result.FirstMessageCreatedAt,
		LastMessageCreatedAt:  result.LastMessageCreatedAt,
	}, err
}
