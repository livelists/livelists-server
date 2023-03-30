package websocket

import (
	"github.com/gobwas/ws/wsutil"
	accessTokenService "github.com/livelists/livelist-server/pkg/services/accessToken"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
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

func (c *WsConnection) publishToSID(payload string) error {
	conn := *c.Connection

	newPaylod := make([]byte, len(payload))
	r := strings.NewReader(payload)

	_, err := io.ReadFull(r, newPaylod)

	if err != nil {
		return err
	}

	err = wsutil.WriteServerMessage(conn, 1, newPaylod)

	if err != nil {
		return err
	}
	return nil
}
