package websocket

import (
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

	cStore.connections[connectionId] = append(cStore.connections[connectionId], newWrapper)
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

type publishToAllSIDsInIdentityArgs struct {
	Identity string
	Payload  []byte
}

func publishToAllSIDsInIdentity(args publishToAllSIDsInIdentityArgs) error {
	connection, ok := cStore.connections[args.Identity]

	if !ok {
		return &connectionNotFoundError{}
	}

	for _, c := range connection {
		c.Connection.publishToSID(args.Payload)
	}

	return nil
}
