package participant

import (
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource"
	"github.com/livelists/livelist-server/pkg/shared"
	"github.com/livelists/livelist-server/pkg/shared/helpers"
)

type LoadParticipantsArgs struct {
	Payload             wsMessages.LoadParticipantsReq
	ChannelIdentifier   string
	RequesterIdentifier string
	WS                  shared.WsRoom
}

func LoadParticipants(args *LoadParticipantsArgs) {
	var participants, err = datasource.GetShortParticipants(datasource.GetShortParticipantsArgs{
		ChannelIdentifier: args.ChannelIdentifier,
		Limit:             1000,
	})

	if err == nil {
		participantsPb := make([]*wsMessages.ParticipantShortInfo, len(participants))

		for i, p := range participants {
			participantsPb[i] = &wsMessages.ParticipantShortInfo{
				Identifier: p.Identifier,
				LastSeenAt: helpers.DateToTimeStamp(p.LastSeenAt),
				IsOnline:   p.IsOnline,
				CustomData: helpers.CustomDataFormat(p.CustomData),
			}
		}

		response := wsMessages.InBoundMessage_LoadParticipantsRes{
			LoadParticipantsRes: &wsMessages.LoadParticipantsRes{
				Participants: participantsPb,
				PageSize:     1000,
			},
		}

		args.WS.PublishMessage(shared.PublishMessageArgs{
			RoomName: args.WS.GetRoomName(shared.GetRoomNameArgs{
				Identifier: args.RequesterIdentifier,
				Type:       wsMessages.WSRoomTypes_Participant,
			}),
			Data: wsMessages.InBoundMessage{Message: &response},
		})
	}
}
