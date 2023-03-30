package shared

type JoinToRoomArgs struct {
	RoomName             string
	WsConnectionIdentity string
}
type PublishMessageArgs struct {
	RoomName string
	Event    string
	Data     map[string]any
}
type WsRoom interface {
	JoinToRoom(args JoinToRoomArgs)
	PublishMessage(args PublishMessageArgs) bool
}
