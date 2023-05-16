package channel

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/services/message"
	"github.com/livelists/livelist-server/pkg/shared"
	"time"
)

type LoadMoreMessagesArgs struct {
	Payload             wsMessages.LoadMoreMessages
	ChannelId           string
	RequesterIdentifier string
	WS                  shared.WsRoom
}

func LoadMoreMessages(args *LoadMoreMessagesArgs) {
	var startFromDate = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

	if args.Payload.FirstLoadedCreatedAt != nil {
		startFromDate = args.Payload.FirstLoadedCreatedAt.AsTime()
	}

	fmt.Println(startFromDate)

	messages, totalCount, err := message.GetMessages(message.GetMessagesArgs{
		StartFromDate:     startFromDate,
		Offset:            int(args.Payload.SkipFromFirstLoaded),
		PageSize:          int(args.Payload.PageSize),
		ChannelIdentifier: args.ChannelId,
	})

	response := wsMessages.InBoundMessage_LoadMoreMessagesRes{
		LoadMoreMessagesRes: &wsMessages.LoadMoreMessagesRes{
			RequestInfo: &wsMessages.LoadMoreMessagesRequestInfo{
				PageSize:             args.Payload.PageSize,
				FirstLoadedCreatedAt: args.Payload.FirstLoadedCreatedAt,
				SkipFromFirstLoaded:  args.Payload.SkipFromFirstLoaded,
			},
			IsSuccess:     true,
			TotalMessages: totalCount,
			Messages:      messages,
		},
	}

	if err != nil {
		response = wsMessages.InBoundMessage_LoadMoreMessagesRes{
			LoadMoreMessagesRes: &wsMessages.LoadMoreMessagesRes{
				RequestInfo: &wsMessages.LoadMoreMessagesRequestInfo{
					PageSize:             args.Payload.PageSize,
					FirstLoadedCreatedAt: args.Payload.FirstLoadedCreatedAt,
					SkipFromFirstLoaded:  args.Payload.SkipFromFirstLoaded,
				},
				IsSuccess:     false,
				TotalMessages: totalCount,
				Messages:      messages,
			},
		}
	}

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.RequesterIdentifier,
			Type:       shared.RoomName_participant,
		}),
		Data: wsMessages.InBoundMessage{Message: &response},
	})
}
