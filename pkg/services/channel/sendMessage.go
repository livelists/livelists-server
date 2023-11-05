package channel

import (
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/services/message"
	"github.com/livelists/livelist-server/pkg/services/participant"
	"github.com/livelists/livelist-server/pkg/shared"
)

type SendMessageArgs struct {
	Payload          wsMessages.SendMessage
	ChannelId        string
	SenderIdentifier string
	WS               shared.WsRoom
}

func SendMessage(args *SendMessageArgs) {
	var createdMessage, messageCreateErr = message.CreateMessage(&message.CreateMessageArgs{
		Payload: message.CreateMessagePayload{
			Text:       args.Payload.Text,
			LocalId:    args.Payload.LocalId,
			CustomData: args.Payload.CustomData,
		},
		ChannelId:        args.ChannelId,
		SenderIdentifier: &args.SenderIdentifier,
		WS:               args.WS,
		Type:             wsMessages.MessageType_ParticipantCreated,
		SubType:          wsMessages.MessageSubType_TextMessage,
	})

	if messageCreateErr == nil {
		participant.UpdateLastMessageSeenAt(&participant.UpdateLastMessageSeenAtArgs{
			ChannelId:           args.ChannelId,
			RequesterIdentifier: args.SenderIdentifier,
			LastSeenAtUnixMS:    createdMessage.CreatedAt.UnixMilli(),
			WS:                  args.WS,
		})
	}
}
