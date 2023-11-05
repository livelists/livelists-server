package message

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
	"github.com/livelists/livelist-server/pkg/shared"
	"github.com/livelists/livelist-server/pkg/shared/helpers"
)

type CreateMessagePayload struct {
	Text       string
	LocalId    string
	CustomData *wsMessages.CustomData
}

type CreateMessageArgs struct {
	Payload          CreateMessagePayload
	ChannelId        string
	SenderIdentifier *string
	Type             wsMessages.MessageType
	SubType          wsMessages.MessageSubType
	WS               shared.WsRoom
}

func CreateMessage(args *CreateMessageArgs) (*mongoSchemes.Message, error) {
	createdMessage, err := datasource.AddMessage(datasource.AddMessageArgs{
		Text:              args.Payload.Text,
		CustomData:        args.Payload.CustomData,
		Type:              args.Type.String(),
		SubType:           args.SubType.String(),
		SenderIdentifier:  args.SenderIdentifier,
		ChannelIdentifier: args.ChannelId,
	})

	if err != nil {
		fmt.Println("createMessageErr", err)
		return nil, err
	}

	message := wsMessages.Message{
		Id:                createdMessage.ID.Hex(),
		Text:              args.Payload.Text,
		ChannelIdentifier: args.ChannelId,
		LocalId:           args.Payload.LocalId,
		Type:              args.Type,
		SubType:           args.SubType,
		CustomData:        args.Payload.CustomData,
		CreatedAt:         helpers.DateToTimeStamp(createdMessage.CreatedAt),
	}

	if args.SenderIdentifier != nil {
		participant, _ := datasource.FindParticipantByIdentifierAndChannel(datasource.FindPByIdAndChannelArgs{
			Identifier: *args.SenderIdentifier,
			ChannelId:  args.ChannelId,
		})

		message.Sender = &wsMessages.ParticipantShortInfo{
			Identifier: *args.SenderIdentifier,
			LastSeenAt: &timestamp.Timestamp{
				Seconds: 0,
				Nanos:   0,
			},
			IsOnline:   true,
			CustomData: helpers.CustomDataFormat(participant.CustomData),
		}
	}

	newM := wsMessages.InBoundMessage_NewMessage{
		NewMessage: &message,
	}

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.ChannelId,
			Type:       wsMessages.WSRoomTypes_Channel,
		}),
		Data: wsMessages.InBoundMessage{Message: &newM},
	})

	return &createdMessage, nil
}
