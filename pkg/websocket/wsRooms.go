package websocket

import (
	"github.com/livelists/livelist-server/pkg/shared"
	"google.golang.org/protobuf/proto"
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

func checkIsIdentityNotAlreadyConnected(roomConnections []string, connection string) bool {
	for _, c := range roomConnections {
		if c == connection {
			return false
		}
	}
	return true
}

func (w WsRoom) JoinToRoom(args shared.JoinToRoomArgs) {
	existedConnections := roomsStore.Rooms[args.RoomName].Connections
	if checkIsIdentityNotAlreadyConnected(existedConnections, args.WsConnectionIdentity) {
		newConnections := append(existedConnections, args.WsConnectionIdentity)
		roomCopy := WsRoom{}
		roomCopy.Connections = newConnections
		roomsStore.Rooms[args.RoomName] = roomCopy
	}
}

func (w WsRoom) PublishMessage(args shared.PublishMessageArgs) bool {
	room, ok := roomsStore.Rooms[args.RoomName]

	if !ok {
		return false
	}
	if len(room.Connections) == 0 {
		return false
	}

	messageBytes, err := proto.Marshal(&args.Data)

	if err != nil {
		return false
	}
	for _, connectionId := range room.Connections {
		publishToAllSIDsInIdentity(publishToAllSIDsInIdentityArgs{
			Identity: connectionId,
			Payload:  messageBytes,
		})
	}

	return true
}
