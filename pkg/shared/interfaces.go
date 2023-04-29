package shared

import "github.com/livelists/livelist-server/contracts/wsMessages"

type RoomName string
type RoomNameType int32

const (
	RoomName_channel     RoomNameType = 0
	RoomName_participant RoomNameType = 1
)

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
	Type       RoomNameType
}
type WsRoom interface {
	JoinToRoom(args JoinToRoomArgs)
	PublishMessage(args PublishMessageArgs) bool
	GetRoomName(args GetRoomNameArgs) RoomName
}
