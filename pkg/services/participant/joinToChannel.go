package participant

import (
	"fmt"
	"github.com/livelists/livelist-server/pkg/shared"
)

type JoinToChannelArgs struct {
	Payload      map[string]any
	WsIdentifier string
	WS           shared.WsRoom
}

func JoinToChannel(args JoinToChannelArgs) {
	channelId := args.Payload["channelId"].(string)
	args.WS.JoinToRoom(shared.JoinToRoomArgs{
		WsConnectionIdentity: args.WsIdentifier,
		RoomName:             channelId,
	})

	fmt.Println(channelId, args.WsIdentifier)
}
