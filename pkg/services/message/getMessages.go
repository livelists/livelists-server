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

func GetMessages(args GetMessagesArgs) ([]*wsMessages.Message, int64, error) {
	messages, totalCount, err := datasource.GetMessagesFromChannel(datasource.GetMessagesFromChannelArgs{
		ChannelIdentifier: args.ChannelIdentifier,
		Limit:             args.PageSize,
		Skip:              args.Offset,
		StartFromDate:     args.StartFromDate,
	})

	return helpers.MongoMessagesToPB(messages), totalCount, err
}
