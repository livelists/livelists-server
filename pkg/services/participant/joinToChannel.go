package participant

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/services/message"
	"github.com/livelists/livelist-server/pkg/shared"
	"github.com/livelists/livelist-server/pkg/shared/helpers"
	"time"
)

type JoinToChannelArgs struct {
	Payload      wsMessages.JoinChannel
	WsIdentifier string
	ChannelId    string
	WS           shared.WsRoom
}

func JoinToChannel(args JoinToChannelArgs) {
	args.WS.JoinToRoom(shared.JoinToRoomArgs{
		WsConnectionIdentity: args.WsIdentifier,
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.ChannelId,
			Type:       shared.RoomName_channel,
		}),
	})
	args.WS.JoinToRoom(shared.JoinToRoomArgs{
		WsConnectionIdentity: args.WsIdentifier,
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.WsIdentifier,
			Type:       shared.RoomName_participant,
		}),
	})

	messages, err := message.GetMessages(message.GetMessagesArgs{
		PageSize:          int(args.Payload.InitialPageSize),
		Offset:            int(args.Payload.InitialOffset),
		ChannelIdentifier: args.ChannelId,
		StartFromDate:     time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
	})

	fmt.Println("messagesLength", len(messages))
	messagesPb := make([]*wsMessages.Message, len(messages))

	for i, m := range messages {
		messagesPb[i] = &wsMessages.Message{
			Id:         m.Id,
			Text:       m.Text,
			SubType:    wsMessages.MessageSubType(wsMessages.MessageSubType_value[m.SubType]),
			Type:       wsMessages.MessageType(wsMessages.MessageSubType_value[m.Type]),
			LocalId:    nil,
			CustomData: helpers.CustomDataFormat(m.CustomData),
			Sender: &wsMessages.ParticipantShortInfo{
				Identifier: m.Participant.Identifier,
				CustomData: helpers.CustomDataFormat(m.Participant.CustomData),
			},
			CreatedAt: helpers.DateToTimeStamp(m.CreatedAt),
		}
	}

	meJoinedMessage := wsMessages.InBoundMessage_MeJoinedToChannel{
		MeJoinedToChannel: &wsMessages.MeJoinedToChannel{
			MeIdentifier: args.WsIdentifier,
			IsSuccess:    true,
			Channel: &wsMessages.ChannelInitialInfo{
				ChannelId:       args.ChannelId,
				HistoryMessages: messagesPb,
			},
		},
	}

	if err != nil {
		meJoinedMessage = wsMessages.InBoundMessage_MeJoinedToChannel{
			MeJoinedToChannel: &wsMessages.MeJoinedToChannel{
				MeIdentifier: args.WsIdentifier,
				IsSuccess:    false,
				Channel: &wsMessages.ChannelInitialInfo{
					ChannelId:       args.ChannelId,
					HistoryMessages: messagesPb,
				},
			},
		}
	}

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.WsIdentifier,
			Type:       shared.RoomName_participant,
		}),
		Data: wsMessages.InBoundMessage{Message: &meJoinedMessage},
	})
}
