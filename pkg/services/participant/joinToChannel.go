package participant

import (
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource"
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

func JoinToChannel(args *JoinToChannelArgs) {
	var meParticipant, err = datasource.FindParticipantByIdentifierAndChannel(datasource.FindPByIdAndChannelArgs{
		ChannelId:  args.ChannelId,
		Identifier: args.WsIdentifier,
	})

	JoinToChannelRoom(&JoinToChannelRoomArgs{
		WsIdentifier: args.WsIdentifier,
		ChannelId:    args.ChannelId,
		WS:           args.WS,
	})

	messages, messagesCount, err := message.GetMessages(message.GetMessagesArgs{
		PageSize:          int(args.Payload.InitialPageSize),
		Offset:            int(args.Payload.InitialOffset),
		ChannelIdentifier: args.ChannelId,
		StartFromDate:     time.Now(),
	})

	meJoinedMessage := wsMessages.InBoundMessage_MeJoinedToChannel{
		MeJoinedToChannel: &wsMessages.MeJoinedToChannel{
			Me: &wsMessages.MeJoined{
				Identifier: args.WsIdentifier,
				Grants: &wsMessages.ChannelParticipantGrants{
					Admin:        &meParticipant.Grants.Admin,
					SendMessage:  &meParticipant.Grants.SendMessage,
					ReadMessages: &meParticipant.Grants.ReadMessages,
				},
				CustomData: helpers.CustomDataFormat(meParticipant.CustomData),
			},
			IsSuccess: true,
			Channel: &wsMessages.ChannelInitialInfo{
				ChannelId:       args.ChannelId,
				HistoryMessages: messages,
				TotalMessages:   messagesCount,
			},
		},
	}

	if err != nil {
		meJoinedMessage = wsMessages.InBoundMessage_MeJoinedToChannel{
			MeJoinedToChannel: &wsMessages.MeJoinedToChannel{
				Me:        &wsMessages.MeJoined{},
				IsSuccess: false,
				Channel: &wsMessages.ChannelInitialInfo{
					ChannelId:       args.ChannelId,
					TotalMessages:   messagesCount,
					HistoryMessages: messages,
				},
			},
		}
	}

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.WsIdentifier,
			Type:       wsMessages.WSRoomTypes_Participant,
		}),
		Data: wsMessages.InBoundMessage{Message: &meJoinedMessage},
	})
}
