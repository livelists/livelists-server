package websocket

import (
	"fmt"
	"github.com/livelists/livelist-server/pkg/shared"
)

func init() {
	roomsStore.Rooms = make(map[string]WsRoom)
}

type WsRoom struct {
	Connections []string
}

type WsRooms struct {
	Rooms map[string]WsRoom
}

var roomsStore WsRooms

func (w WsRoom) JoinToRoom(args shared.JoinToRoomArgs) {
	newConnections := append(roomsStore.Rooms[args.RoomName].Connections, args.WsConnectionIdentity)
	roomCopy := WsRoom{}
	roomCopy.Connections = newConnections
	roomsStore.Rooms[args.RoomName] = roomCopy
}

func (w WsRoom) PublishMessage(args shared.PublishMessageArgs) bool {
	room, ok := roomsStore.Rooms[args.RoomName]

	if !ok {
		return false
	}
	if len(room.Connections) == 0 {
		return false
	}

	for _, connectionId := range room.Connections {
		publishToAllSIDsInIdentity(publishToAllSIDsInIdentityArgs{
			Identity: connectionId,
			Payload:  args.Data["text"].(string),
		})
		fmt.Println(room.Connections)
		fmt.Println(&room.Connections)
		fmt.Println("publish to connection", connectionId)
	}

	return true
}
