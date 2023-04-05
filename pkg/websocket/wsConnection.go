package websocket

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	accessTokenService "github.com/livelists/livelist-server/pkg/services/accessToken"
	"net"
	"net/http"
	"net/url"
)

type WsConnection struct {
	AccessToken accessTokenService.AccessToken
	Connection  *net.Conn
}

type TokenParseError struct{}

func (m *TokenParseError) Error() string {
	return "access token param not found"
}

func (c *WsConnection) onAuth(r *http.Request) (bool, error) {
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

	c.AccessToken = accessToken

	return true, nil
}

func (c *WsConnection) addConnection(conn *net.Conn) {
	c.Connection = conn
}

func (c *WsConnection) publishToSID(payload []byte) error {
	conn := *c.Connection

	err := wsutil.WriteServerMessage(conn, ws.OpBinary, payload)

	if err != nil {
		return err
	}
	return nil
}
