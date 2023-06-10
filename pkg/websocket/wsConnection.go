package websocket

import (
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	accessTokenService "github.com/livelists/livelist-server/pkg/services/accessToken"
	"github.com/livelists/livelist-server/pkg/services/participant"
	"net"
	"net/http"
	"net/url"
	"time"
)

type WsConnection struct {
	AccessToken        accessTokenService.AccessToken
	DisconnectCallback func()
	Connection         *net.Conn
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

	participant.UpdateLastSeenAt(&participant.UpdateLastSeenAtArgs{
		WsIdentifier:      WsC.AccessToken.Identifier(),
		ChannelIdentifier: WsC.AccessToken.ChannelId(),
		IsOnline:          true,
		WS:                newWs,
	})
	WsC.Connection = conn
	//WsC.startDetectDisconnection(conn)
}

func (WsC *WsConnection) startDetectDisconnection(conn *net.Conn) {
	one := make([]byte, 1)
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	go func() {
		var c = *conn
		c.SetReadDeadline(time.Now().Add(10 * time.Second))

		for {
			select {
			case <-ticker.C:
				if _, err := c.Read(one); err != nil {
					ticker.Stop()
					fmt.Print("disconnect err", err)
					//WsC.disconnectCallback()
				}
				c.SetReadDeadline(time.Now().Add(10 * time.Second))
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
