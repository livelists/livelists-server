package channel

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/services/participant"
	"github.com/livelists/livelist-server/pkg/shared"
	"github.com/livelists/livelist-server/pkg/shared/helpers"
)

type GetChannelsArgs struct {
	MessagesLimit       int32
	RequesterIdentifier string
	WS                  shared.WsRoom
}

func GetMyChannelsWithMsg(args *GetChannelsArgs) {
	channels, err := loadChannels(&loadChannelsArgs{
		MessagesLimit:       args.MessagesLimit,
		RequesterIdentifier: args.RequesterIdentifier,
	})

	if err != nil {
		fmt.Println("GetMyChannelsError")
		return
	}

	for _, c := range channels {
		participant.JoinToChannelRoom(&participant.JoinToChannelRoomArgs{
			ChannelId:    c.Channel.Identifier,
			WsIdentifier: args.RequesterIdentifier,
			WS:           args.WS,
		})
	}
	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.RequesterIdentifier,
			Type:       wsMessages.WSRoomTypes_Participant,
		}),
		Data: wsMessages.InBoundMessage{
			Message: &wsMessages.InBoundMessage_LoadChannelsWithMsgRes{
				LoadChannelsWithMsgRes: &wsMessages.LoadChannelsWithMsgRes{
					Channels: channels,
				},
			},
		},
	})
}

type loadChannelsArgs struct {
	MessagesLimit       int32
	RequesterIdentifier string
}

func loadChannels(args *loadChannelsArgs) ([]*wsMessages.ChannelWithMsg, error) {
	channels, err := datasource.GetParticipantsChannelsWithMessages(datasource.GetParticipantChWithMsgArgs{
		MessagesLimit:         args.MessagesLimit,
		ParticipantIdentifier: args.RequesterIdentifier,
	})

	channelsPb := make([]*wsMessages.ChannelWithMsg, len(channels))

	for i, ch := range channels {
		channelsPb[i] = &wsMessages.ChannelWithMsg{
			Channel: &wsMessages.ShortChannel{
				Id:         ch.Channel.Id,
				Identifier: ch.Channel.Identifier,
				CustomData: helpers.CustomDataFormat(ch.Channel.CustomData),
			},
			Messages: helpers.MongoMessagesToPB(ch.Messages),
		}
	}

	return channelsPb, err
}
