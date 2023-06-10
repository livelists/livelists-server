package shared

import "github.com/livelists/livelist-server/contracts/wsMessages"

type RoomName string

type JoinToRoomArgs struct {
	RoomName             RoomName
	WsConnectionIdentity string
}
type PublishMessageArgs struct {
	RoomName RoomName
	Data     wsMessages.InBoundMessage
}
type GetRoomNameArgs struct {
	Identifier string
	Type       wsMessages.WSRoomTypes
}
type WsRoom interface {
	JoinToRoom(args JoinToRoomArgs)
	PublishMessage(args PublishMessageArgs) bool
	GetRoomName(args GetRoomNameArgs) RoomName
}
