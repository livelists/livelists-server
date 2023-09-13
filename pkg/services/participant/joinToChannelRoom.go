package participant

import (
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/shared"
)

type JoinToChannelRoomArgs struct {
	WsIdentifier string
	ChannelId    string
	WS           shared.WsRoom
}

func JoinToChannelRoom(args *JoinToChannelRoomArgs) {
	args.WS.JoinToRoom(shared.JoinToRoomArgs{
		WsConnectionIdentity: args.WsIdentifier,
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.ChannelId,
			Type:       wsMessages.WSRoomTypes_Channel,
		}),
	})
}
