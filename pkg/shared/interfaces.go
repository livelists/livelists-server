package shared

import "github.com/livelists/livelist-server/contracts/wsMessages"

type JoinToRoomArgs struct {
	RoomName             string
	WsConnectionIdentity string
}
type PublishMessageArgs struct {
	RoomName string
	Data     wsMessages.InBoundMessage
}
type WsRoom interface {
	JoinToRoom(args JoinToRoomArgs)
	PublishMessage(args PublishMessageArgs) bool
}
