package participant

import (
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/shared"
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
		RoomName:             args.ChannelId,
	})
}
