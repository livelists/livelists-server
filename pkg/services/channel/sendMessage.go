package channel

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/shared"
)

type SendMessageArgs struct {
	Payload          wsMessages.SendMessage
	ChannelId        string
	SenderIdentifier string
	WS               shared.WsRoom
}

func SendMessage(args *SendMessageArgs) {
	createdMessage, err := datasource.AddMessage(datasource.AddMessageArgs{
		Text:              args.Payload.Text,
		CustomData:        args.Payload.CustomData,
		Type:              wsMessages.MessageType_ParticipantCreated.String(),
		SubType:           wsMessages.MessageSubType_TextMessage.String(),
		SenderIdentifier:  args.SenderIdentifier,
		ChannelIdentifier: args.ChannelId,
	})

	if err != nil {
		fmt.Println("createMessageErr", err)
		return
	}

	message := wsMessages.Message{
		Id:   createdMessage.ID.Hex(),
		Text: args.Payload.Text,
		Sender: &wsMessages.ParticipantShortInfo{
			Identifier: args.SenderIdentifier,
			CustomData: args.Payload.CustomData,
		},
		LocalId:    args.Payload.LocalId,
		Type:       wsMessages.MessageType_ParticipantCreated,
		SubType:    wsMessages.MessageSubType_TextMessage,
		CustomData: args.Payload.CustomData,
		CreatedAt: &timestamp.Timestamp{
			Seconds: createdMessage.CreatedAt.Unix(),
			Nanos:   int32(createdMessage.CreatedAt.UnixNano()),
		},
	}

	newM := wsMessages.InBoundMessage_NewMessage{
		NewMessage: &message,
	}

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.ChannelId,
			Type:       shared.RoomName_channel,
		}),
		Data: wsMessages.InBoundMessage{Message: &newM},
	})
}
