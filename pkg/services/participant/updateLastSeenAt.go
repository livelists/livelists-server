package participant

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/shared"
	"time"
)

type UpdateLastSeenAtArgs struct {
	ChannelIdentifier string
	WsIdentifier      string
	IsOnline          bool
	WS                shared.WsRoom
}

func UpdateLastSeenAt(args *UpdateLastSeenAtArgs) {
	var now = time.Now()

	datasource.UpdateParticipantLastSeenAt(datasource.UpdateParticipantLastSeenAtArgs{
		ChannelIdentifier: args.ChannelIdentifier,
		Identifier:        args.WsIdentifier,
		LastSeenAt:        now,
		IsOnline:          args.IsOnline,
	})

	if args.IsOnline {
		publishBecameOnline(publishBecameOnlineArgs{
			WS:                args.WS,
			WsIdentifier:      args.WsIdentifier,
			ChannelIdentifier: args.ChannelIdentifier,
		})
	} else {
		publishBecameOffline(publishBecameOfflineArgs{
			WS:                args.WS,
			WsIdentifier:      args.WsIdentifier,
			ChannelIdentifier: args.ChannelIdentifier,
			LastSeenAt:        now,
		})
	}
}

type publishBecameOfflineArgs struct {
	WS                shared.WsRoom
	ChannelIdentifier string
	WsIdentifier      string
	LastSeenAt        time.Time
}

func publishBecameOffline(args publishBecameOfflineArgs) {
	message := wsMessages.ParticipantBecameOffline{
		Identifier: args.WsIdentifier,
		LastSeenAt: &timestamp.Timestamp{Seconds: int64(args.LastSeenAt.Second())},
	}

	messageWr := wsMessages.InBoundMessage_ParticipantBecameOffline{
		ParticipantBecameOffline: &message,
	}

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.ChannelIdentifier,
			Type:       wsMessages.WSRoomTypes_Channel,
		}),
		Data: wsMessages.InBoundMessage{Message: &messageWr},
	})
}

type publishBecameOnlineArgs struct {
	WS                shared.WsRoom
	ChannelIdentifier string
	WsIdentifier      string
}

func publishBecameOnline(args publishBecameOnlineArgs) {
	message := wsMessages.ParticipantBecameOnline{
		Identifier: args.WsIdentifier,
	}

	messageWr := wsMessages.InBoundMessage_ParticipantBecameOnline{
		ParticipantBecameOnline: &message,
	}

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.ChannelIdentifier,
			Type:       wsMessages.WSRoomTypes_Channel,
		}),
		Data: wsMessages.InBoundMessage{Message: &messageWr},
	})
}
