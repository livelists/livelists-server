package channel

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/services/message"
	"github.com/livelists/livelist-server/pkg/shared"
	"github.com/livelists/livelist-server/pkg/shared/helpers"
	"time"
)

type LoadMoreMessagesArgs struct {
	Payload             wsMessages.LoadMoreMessages
	ChannelId           string
	RequesterIdentifier string
	IsLoadOlder         bool
	WS                  shared.WsRoom
}

func LoadMoreMessages(args *LoadMoreMessagesArgs) {
	var startFromDate = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

	if args.Payload.FirstLoadedCreatedAt != nil {
		startFromDate = args.Payload.FirstLoadedCreatedAt.AsTime()
	}

	fmt.Println(startFromDate)

	messagesResult, err := message.GetMessages(message.GetMessagesArgs{
		StartFromDate:     startFromDate,
		IsLoadOlder:       true,
		Offset:            int(args.Payload.SkipFromFirstLoaded),
		PageSize:          int(args.Payload.PageSize),
		ChannelIdentifier: args.ChannelId,
	})

	fmt.Println(len(messagesResult.Messages))

	response := wsMessages.InBoundMessage_LoadMoreMessagesRes{
		LoadMoreMessagesRes: &wsMessages.LoadMoreMessagesRes{
			RequestInfo: &wsMessages.LoadMoreMessagesRequestInfo{
				PageSize:             args.Payload.PageSize,
				FirstLoadedCreatedAt: args.Payload.FirstLoadedCreatedAt,
				SkipFromFirstLoaded:  args.Payload.SkipFromFirstLoaded,
			},
			IsSuccess:             true,
			FirstMessageCreatedAt: helpers.DateToTimeStamp(messagesResult.FirstMessageCreatedAt),
			LastMessageCreatedAt:  helpers.DateToTimeStamp(messagesResult.LastMessageCreatedAt),
			TotalMessages:         messagesResult.TotalCount,
			Messages:              messagesResult.Messages,
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
				IsSuccess:             false,
				TotalMessages:         0,
				FirstMessageCreatedAt: nil,
				LastMessageCreatedAt:  nil,
				Messages:              []*wsMessages.Message{},
			},
		}
	}

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.RequesterIdentifier,
			Type:       wsMessages.WSRoomTypes_Participant,
		}),
		Data: wsMessages.InBoundMessage{Message: &response},
	})
}
