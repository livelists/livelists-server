package websocket

import (
	"github.com/livelists/livelist-server/pkg/services/participant"
	"golang.org/x/exp/slices"
)

func init() {
	cStore.connections = make(map[string][]*WsConnection)
}

type ConnectionsStore struct {
	connections map[string][]*WsConnection
}

var cStore ConnectionsStore

func AddWsConnection(conn *WsConnection) {
	connectionId := conn.AccessToken.Identifier()

	cStore.connections[connectionId] = append(cStore.connections[connectionId], conn)

	conn.OnDisconnected(func() {
		onWsDisconnected(conn, connectionId)
	})
}

func onWsDisconnected(conn *WsConnection, connectionId string) {
	var disconnectedChannelId = conn.AccessToken.ChannelId()

	if !isHasAnotherConnectionToChannel(IsHasAnotherCh{
		DisconnectedChannelId: disconnectedChannelId,
		ConnectionId:          connectionId,
		Conn:                  conn,
	}) {
		newWs := &WsRoom{}

		participant.UpdateLastSeenAt(&participant.UpdateLastSeenAtArgs{
			WsIdentifier:      conn.AccessToken.Identifier(),
			ChannelIdentifier: conn.AccessToken.ChannelId(),
			IsOnline:          false,
			WS:                newWs,
		})
	}

	var cleanedConnections = removeWsConnection(RemoveWsConnFromArrArgs{
		ConnSidToRemove: conn.Sid,
		Arr:             cStore.connections[connectionId],
	})

	cStore.connections[connectionId] = cleanedConnections
}

type IsHasAnotherCh struct {
	Conn                  *WsConnection
	ConnectionId          string
	DisconnectedChannelId string
}

func isHasAnotherConnectionToChannel(args IsHasAnotherCh) bool {
	var anotherConnectionId = slices.IndexFunc(cStore.connections[args.ConnectionId], func(c *WsConnection) bool {
		return c.AccessToken.ChannelId() == args.DisconnectedChannelId && c.Sid != args.Conn.Sid
	})

	if anotherConnectionId == -1 {
		return false
	}

	return true
}

type RemoveWsConnFromArrArgs struct {
	Arr             []*WsConnection
	ConnSidToRemove string
}

func removeWsConnection(args RemoveWsConnFromArrArgs) []*WsConnection {
	var arr = args.Arr

	var connectionIndex = slices.IndexFunc(args.Arr, func(c *WsConnection) bool {
		return c.Sid == args.ConnSidToRemove
	})

	if connectionIndex == -1 {
		return arr
	}

	arr[connectionIndex] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func GetWSConnections() map[string][]*WsConnection {
	if cStore.connections == nil {
		cStore.connections = make(map[string][]*WsConnection)
	}

	return cStore.connections
}

type connectionNotFoundError struct{}

func (e *connectionNotFoundError) Error() string {
	return "access token param not found"
}

type PublishToAllSIDsInIdentityArgs struct {
	Identity string
	Payload  []byte
}

func PublishToAllSIDsInIdentity(args PublishToAllSIDsInIdentityArgs) error {
	connection, ok := cStore.connections[args.Identity]

	if !ok {
		return &connectionNotFoundError{}
	}

	for _, c := range connection {
		c.PublishToSID(args.Payload)
	}

	return nil
}
