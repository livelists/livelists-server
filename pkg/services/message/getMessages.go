package message

import (
	"github.com/golang/protobuf/ptypes/timestamp"
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

	messagesPb := make([]*wsMessages.Message, len(messages))

	for i, m := range messages {
		messagesPb[i] = &wsMessages.Message{
			Id:         m.Id,
			Text:       m.Text,
			SubType:    wsMessages.MessageSubType(wsMessages.MessageSubType_value[m.SubType]),
			Type:       wsMessages.MessageType(wsMessages.MessageSubType_value[m.Type]),
			LocalId:    "",
			CustomData: helpers.CustomDataFormat(m.CustomData),
			Sender: &wsMessages.ParticipantShortInfo{
				Identifier: m.Participant.Identifier,
				LastSeenAt: &timestamp.Timestamp{
					Seconds: 0,
					Nanos:   0,
				},
				IsOnline:   true,
				CustomData: helpers.CustomDataFormat(m.Participant.CustomData),
			},
			CreatedAt: helpers.DateToTimeStamp(m.CreatedAt),
		}
	}

	return messagesPb, totalCount, err
}
