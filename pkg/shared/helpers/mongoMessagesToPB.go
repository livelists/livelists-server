package helpers

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/datasource/mongoSchemes"
)

func MongoMessagesToPB(messages []mongoSchemes.MessageWithParticipant) []*wsMessages.Message {
	messagesPb := make([]*wsMessages.Message, len(messages))

	for i, m := range messages {
		messagesPb[i] = &wsMessages.Message{
			Id:                m.Id,
			Text:              m.Text,
			SubType:           wsMessages.MessageSubType(wsMessages.MessageSubType_value[m.SubType]),
			Type:              wsMessages.MessageType(wsMessages.MessageType_value[m.Type]),
			LocalId:           "",
			ChannelIdentifier: m.ChannelIdentifier,
			CustomData:        CustomDataFormat(m.CustomData),
			Sender: &wsMessages.ParticipantShortInfo{
				Identifier: m.Participant.Identifier,
				LastSeenAt: &timestamp.Timestamp{
					Seconds: 0,
					Nanos:   0,
				},
				IsOnline:   true,
				CustomData: CustomDataFormat(m.Participant.CustomData),
			},
			CreatedAt: DateToTimeStamp(m.CreatedAt),
		}
	}

	return messagesPb
}
