package participant

import "github.com/livelists/livelist-server/pkg/shared"

type SendMessageArgs struct {
	Text      string
	ChannelId string
	WS        shared.WsRoom
}

func SendMessage(args SendMessageArgs) {
	args.WS.PublishMessage(shared.PublishMessageArgs{
		RoomName: args.ChannelId,
		Event:    "channel:message:new",
		Data: map[string]any{
			"text": args.Text,
		},
	})
}
