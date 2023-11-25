package websocket

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/livelists/livelist-server/contracts/wsMessages"
	accessTokenService "github.com/livelists/livelist-server/pkg/services/accessToken"
	"github.com/livelists/livelist-server/pkg/services/participant"
	"github.com/livelists/livelist-server/pkg/shared"
	"github.com/livelists/livelist-server/pkg/shared/helpers"
	"net"
	"net/http"
	"net/url"
	"time"
)

type WsConnection struct {
	Sid                string
	AccessToken        accessTokenService.AccessToken
	DisconnectCallback func()
	Connection         *net.Conn
}

func NewWsConnection() *WsConnection {
	return &WsConnection{
		Sid: helpers.RandStringRunes(8),
	}
}

type TokenParseError struct{}

func (m *TokenParseError) Error() string {
	return "access token param not found"
}

func (WsC *WsConnection) OnAuth(r *http.Request) (bool, error) {
	parsedUrl, err := url.Parse(r.RequestURI)
	if err != nil {
		return false, err
	}

	accessTokenArr := parsedUrl.Query()["accessToken"]

	if len(accessTokenArr) == 0 {
		return false, &TokenParseError{}
	}

	accessTokenStr := accessTokenArr[0]
	accessToken := accessTokenService.AccessToken{}

	accessToken.Parse(accessTokenStr)

	WsC.AccessToken = accessToken

	return true, nil
}

func (WsC *WsConnection) OnDisconnected(callback func()) {
	WsC.DisconnectCallback = callback
}

func (WsC *WsConnection) AddConnection(conn *net.Conn) {
	newWs := &WsRoom{}

	newWs.JoinToRoom(shared.JoinToRoomArgs{
		WsConnectionIdentity: WsC.AccessToken.Identifier(),
		RoomName: newWs.GetRoomName(shared.GetRoomNameArgs{
			Identifier: WsC.AccessToken.Identifier(),
			Type:       wsMessages.WSRoomTypes_Participant,
		}),
	})

	participant.UpdateLastSeenAt(&participant.UpdateLastSeenAtArgs{
		WsIdentifier: WsC.AccessToken.Identifier(),
		IsOnline:     true,
		WS:           newWs,
	})

	WsC.Connection = conn

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				conn := *WsC.Connection
				wsutil.WriteServerMessage(conn, ws.OpPing, ws.CompiledPing)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (WsC *WsConnection) PublishToSID(payload []byte) error {
	conn := *WsC.Connection

	err := wsutil.WriteServerMessage(conn, ws.OpBinary, payload)

	if err != nil {
		return err
	}
	return nil
}
