package participant

import (
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/shared"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UpdateLastMessageSeenAtArgs struct {
	ChannelId           string
	RequesterIdentifier string
	LastSeenAt          *timestamppb.Timestamp
	WS                  shared.WsRoom
}

func UpdateLastMessageSeenAt(args *UpdateLastMessageSeenAtArgs) {
	var meParticipant, err = datasource.FindParticipantByIdentifierAndChannel(datasource.FindPByIdAndChannelArgs{
		ChannelId:  args.ChannelId,
		Identifier: args.RequesterIdentifier,
	})

	if err != nil {
		return
	}

	var isAfterCurrentDate = args.LastSeenAt.GetSeconds() > meParticipant.LastSeenMessageCreatedAt.Unix()

	if !isAfterCurrentDate {
		return
	}

	datasource.UpdateLastMessageSeenAt(datasource.UpdateLastMessageSeenAtArgs{
		ChannelIdentifier: args.ChannelId,
		Identifier:        args.RequesterIdentifier,
		LastSeenAt:        args.LastSeenAt,
	})

	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
			Identifier: args.RequesterIdentifier,
			Type:       wsMessages.WSRoomTypes_Participant,
		}),
		Data: wsMessages.InBoundMessage{Message: &wsMessages.InBoundMessage_UpdateLastSeenMessageAtRes{
			UpdateLastSeenMessageAtRes: &wsMessages.UpdateLastSeenMessageAtRes{
				ChannelId:  args.ChannelId,
				LastSeenAt: args.LastSeenAt,
			},
		}},
	})
}
