package websocket

import (
	"fmt"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	"github.com/livelists/livelist-server/pkg/services/channel"
	"github.com/livelists/livelist-server/pkg/services/participant"
	"google.golang.org/protobuf/proto"
)

type SocketEvent struct {
	Event string         `json:"event"`
	Room  map[string]any `json:"room"`
	Data  map[string]any `json:"data"`
}

func HandleEvent(conn *WsConnection, message []byte) error {
	parsedMessage := wsMessages.OutBoundMessage{}

	err := proto.Unmarshal(message, &parsedMessage)

	if err != nil {
		fmt.Println("Event parse error", err)
		return err
	}

	newWs := WsRoom{}
	switch parsedMessage.Message.(type) {
	case *wsMessages.OutBoundMessage_JoinChannel:
		chatJoin := parsedMessage.GetJoinChannel()
		participant.JoinToChannel(participant.JoinToChannelArgs{
			Payload:      *chatJoin,
			WsIdentifier: conn.AccessToken.Identifier(),
			ChannelId:    conn.AccessToken.ChannelId(),
			WS:           newWs,
		})
	case *wsMessages.OutBoundMessage_SendMessage:
		sendMessage := parsedMessage.GetSendMessage()
		channel.SendMessage(channel.SendMessageArgs{
			Payload:          *sendMessage,
			ChannelId:        conn.AccessToken.ChannelId(),
			SenderIdentifier: conn.AccessToken.Identifier(),
			WS:               newWs,
		})
	}

	return nil
}
