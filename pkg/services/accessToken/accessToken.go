package accessToken

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/livelists/livelist-server/pkg/config"
	"github.com/livelists/livelist-server/pkg/shared/helpers"
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

func (at *AccessToken) Parse(tokenStr string) (bool, error) {
	configData := config.GetConfig()

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configData.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sendMessage := false
		admin := false
		readMessages := false
		identifier := ""

		if claims["SendMessage"] != nil {
			sendMessage = claims["SendMessage"].(bool)
		}

		if claims["Admin"] != nil {
			admin = claims["Admin"].(bool)
		}

		if claims["ReadMessages"] != nil {
			readMessages = claims["ReadMessages"].(bool)
		}

		if claims["Identifier"] != nil {
			identifier = claims["Identifier"].(string)
		}

		if claims["isServiceRoot"] != nil {
			isRoot := claims["isServiceRoot"].(bool)
			at.isServiceRoot = &isRoot
		}
		at.AddGrants(GrantsData{
			SendMessage:  &sendMessage,
			Admin:        &admin,
			ReadMessages: &readMessages,
		})

		at.AddUser(identifier)
		at.isValid = &token.Valid
	} else {
		return false, err
	}

	return true, nil
}

func (at *AccessToken) Sign() (string, error) {
	configData := config.GetConfig()

	now := time.Now()
	expirationDate := now.Add(tokenExpiration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"IsServiceRoot": helpers.FalseIfNil(at.isServiceRoot),
		"Identifier":    at.identifier,
		"SendMessage":   helpers.FalseIfNil(at.grants.SendMessage),
		"ReadMessages":  helpers.FalseIfNil(at.grants.ReadMessages),
		"Admin":         helpers.FalseIfNil(at.grants.Admin),
		"exp":           expirationDate,
	})

	tokenString, err := token.SignedString([]byte(configData.SecretKey))

	return tokenString, err
}

func (at *AccessToken) IsServiceRoot() bool {
	return helpers.FalseIfNil(at.isServiceRoot)
}

func (at *AccessToken) Grants() GrantsData {
	return at.grants
}

func (at *AccessToken) Identifier() string {
	return at.identifier
}
func (at *AccessToken) IsValid() bool {
	return helpers.FalseIfNil(at.isValid)
}
