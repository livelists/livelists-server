package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/livelists/livelist-server/pkg/services/participant"
)

const CHANNEL_JOIN_EVENT = "channel:join"
const CHANNEL_SEND_MESSAGE = "channel:send:message"

type SocketEvent struct {
	Event string         `json:"event"`
	Room  map[string]any `json:"room"`
	Data  map[string]any `json:"data"`
}

func HandleEvent(conn *WsConnection, event []byte) error {
	parsedEvent := SocketEvent{}
	err := json.Unmarshal(event, &parsedEvent)

	if err != nil {
		fmt.Println("Event parse error", err)
		return err
	}

	newWs := WsRoom{}
	switch parsedEvent.Event {
	case CHANNEL_JOIN_EVENT:
		participant.JoinToChannel(participant.JoinToChannelArgs{
			Payload:      parsedEvent.Room,
			WsIdentifier: conn.AccessToken.Identifier(),
			WS:           newWs,
		})
	case CHANNEL_SEND_MESSAGE:
		participant.SendMessage(participant.SendMessageArgs{
			Text:      parsedEvent.Data["text"].(string),
			ChannelId: parsedEvent.Room["channelId"].(string),
			WS:        newWs,
		})
	}

	return nil
}
