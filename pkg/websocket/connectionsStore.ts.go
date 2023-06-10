package websocket

import (
	"fmt"
	"github.com/livelists/livelist-server/pkg/shared/helpers"
)

func init() {
	cStore.connections = make(map[string][]ConnectionWrapper)
}

type ConnectionWrapper struct {
	Sid        string
	Connection *WsConnection
}

type ConnectionsStore struct {
	connections map[string][]ConnectionWrapper
}

var cStore ConnectionsStore

func AddWsConnection(conn *WsConnection) {
	connectionId := conn.AccessToken.Identifier()
	newWrapper := ConnectionWrapper{
		Sid:        helpers.RandStringRunes(8),
		Connection: conn,
	}

	conn.OnDisconnected(func() {
		fmt.Println(len(cStore.connections[connectionId]))

		fmt.Print("On disconnected")
	})

	cStore.connections[connectionId] = append(cStore.connections[connectionId], newWrapper)
}

func removeWsConnection(arr []ConnectionWrapper, i int) []ConnectionWrapper {
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func GetWSConnections() map[string][]ConnectionWrapper {
	if cStore.connections == nil {
		cStore.connections = make(map[string][]ConnectionWrapper)
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
		c.Connection.PublishToSID(args.Payload)
	}

	return nil
}
