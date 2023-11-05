package websocket

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/services/channel"
	"github.com/livelists/livelist-server/pkg/services/customEvents"
	"github.com/livelists/livelist-server/pkg/services/participant"
	"google.golang.org/protobuf/proto"
)

type SocketEvent struct {
	Event string         `json:"event"`
	Room  map[string]any `json:"room"`
	Data  map[string]any `json:"data"`
}

func HandleEvent(conn *WsConnection, message []byte, wsRoom *WsRoom) error {
	parsedMessage := wsMessages.OutBoundMessage{}

	err := proto.Unmarshal(message, &parsedMessage)

	if err != nil {
		fmt.Println("Event parse error", err)
		return err
	}

	switch parsedMessage.Message.(type) {
	case *wsMessages.OutBoundMessage_JoinChannel:
		channelJoin := parsedMessage.GetJoinChannel()
		participant.JoinToChannel(&participant.JoinToChannelArgs{
			Payload:      *channelJoin,
			WsIdentifier: conn.AccessToken.Identifier(),
			ChannelId:    channelJoin.ChannelId,
			WS:           wsRoom,
		})
	case *wsMessages.OutBoundMessage_SendMessage:
		sendMessage := parsedMessage.GetSendMessage()
		channel.SendMessage(&channel.SendMessageArgs{
			Payload:          *sendMessage,
			ChannelId:        sendMessage.ChannelId,
			SenderIdentifier: conn.AccessToken.Identifier(),
			WS:               wsRoom,
		})
	case *wsMessages.OutBoundMessage_LoadMoreMessages:
		payload := parsedMessage.GetLoadMoreMessages()

		channel.LoadMoreMessages(&channel.LoadMoreMessagesArgs{
			Payload:             *payload,
			ChannelId:           payload.ChannelId,
			RequesterIdentifier: conn.AccessToken.Identifier(),
			WS:                  wsRoom,
		})
	case *wsMessages.OutBoundMessage_LoadParticipantsReq:
		payload := parsedMessage.GetLoadParticipantsReq()
		participant.LoadParticipants(&participant.LoadParticipantsArgs{
			Payload:             *payload,
			ChannelIdentifier:   payload.ChannelId,
			RequesterIdentifier: conn.AccessToken.Identifier(),
			WS:                  wsRoom,
		})
	case *wsMessages.OutBoundMessage_SendCustomEvent:
		payload := parsedMessage.GetSendCustomEvent()
		customEvents.SendEvent(&customEvents.SendEventArgs{
			Payload: *payload,
			WS:      wsRoom,
		})
	case *wsMessages.OutBoundMessage_LoadChannelsWithMsgReq:
		payload := parsedMessage.GetLoadChannelsWithMsgReq()
		channel.GetMyChannelsWithMsg(&channel.GetChannelsArgs{
			RequesterIdentifier: conn.AccessToken.Identifier(),
			MessagesLimit:       payload.MessagesLimit,
			WS:                  wsRoom,
		})
	case *wsMessages.OutBoundMessage_UpdateLastSeenMessageAtReq:
		payload := parsedMessage.GetUpdateLastSeenMessageAtReq()
		participant.UpdateLastMessageSeenAt(&participant.UpdateLastMessageSeenAtArgs{
			ChannelId:           payload.ChannelId,
			RequesterIdentifier: conn.AccessToken.Identifier(),
			LastSeenAtUnixMS:    payload.LastSeenAtUnixMS,
			WS:                  wsRoom,
		})
	}

	return nil
}
