package accessToken

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/shared"
	"time"
)

const tokenExpiration = 24 * time.Hour

type GrantsData struct {
	SendMessage  *bool
	ReadMessages *bool
	Admin        *bool
}

type AccessToken struct {
	isServiceRoot *bool
	grants        GrantsData
	identifier    string
	isValid       *bool
}

func (at *AccessToken) AddGrants(grants GrantsData) {
	at.grants = grants
}

func (at *AccessToken) AddUser(identifier string) {
	at.identifier = identifier
}

func (at *AccessToken) Parse(token string) string {
	return token
}

func (at *AccessToken) Sign() (string, error) {
	configData := config.GetConfig()

	now := time.Now()
	expirationDate := now.Add(tokenExpiration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"IsServiceRoot": shared.FalseIfNil(at.isServiceRoot),
		"Identifier":    at.identifier,
		"SendMessage":   shared.FalseIfNil(at.grants.SendMessage),
		"ReadMessages":  shared.FalseIfNil(at.grants.ReadMessages),
		"Admin":         shared.FalseIfNil(at.grants.Admin),
		"exp":           expirationDate,
	})

	tokenString, err := token.SignedString(configData.SecretKey)

	return tokenString, err
}

func (at *AccessToken) IsServiceRoot() bool {
	return shared.FalseIfNil(at.isServiceRoot)
}

func (at *AccessToken) Grants() GrantsData {
	return at.grants
}

func (at *AccessToken) Identifier() string {
	return at.identifier
}
func (at *AccessToken) IsValid() bool {
	return shared.FalseIfNil(at.isValid)
}
