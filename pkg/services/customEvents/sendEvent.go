package customEvents

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/shared"
)

type SendEventArgs struct {
	Payload wsMessages.CustomEvent
	WS      shared.WsRoom
}

func SendEvent(args *SendEventArgs) {
	newCustomEvent := wsMessages.InBoundMessage_NewCustomEvent{
		NewCustomEvent: &wsMessages.CustomEvent{
			RoomType:       args.Payload.RoomType,
			RoomIdentifier: args.Payload.RoomIdentifier,
			EventName:      args.Payload.EventName,
			CustomData:     args.Payload.CustomData,
		},
	}

	fmt.Print(args.Payload.RoomIdentifier)

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.Payload.RoomIdentifier,
			Type:       args.Payload.RoomType,
		}),
		Data: wsMessages.InBoundMessage{Message: &newCustomEvent},
	})
}
