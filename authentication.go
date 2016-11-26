package gofycat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GrantType int

var grant = map[GrantType]string{Client: "client_credentials", Password: "password"}

const (
	ILLEGAL GrantType = iota
	Client
	Password
)

type Cat struct {
	auth                   *authentication
	ClientID, ClientSecret string
	Type                   GrantType
}

type authentication struct {
	TokenType          string `json:"token_type"`
	Scope              string `json:"scope"`
	ExpiresIn          int    `json:"expires_in"`
	Token              string `json:"access_token"`
	RefreshTokenExpiry int    `json:"refresh_token_expires_in,omitempty"`
	RefreshToken       string `json:"refresh_token,omitempty"`
	ResourceOwner      string `json:"resource_owner,omitempty"`
}

func unmarshal(r *http.Response, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func New(clientid, secret string, t GrantType) *Cat {
	c := &Cat{
		ClientID:     clientid,
		ClientSecret: secret,
		Type:         t,
		auth: &authentication{},
	}
	
	return c
}

func (g *Cat) authenticate(grantType GrantType, args ...string) (*authentication, error) {
	switch grantType {
	case Client:
		return clientCredentials(grant[grantType], args...)
	case Password:
		return passwordGrant(grant[grantType], args...)
	default:
		return nil, fmt.Errorf("Invalid grant type '%s'", grant[grantType])
	}
}

func clientCredentials(gt string, args ...string) (*authentication, error) {
	a := &authentication{}
	payload := struct {
		GrantType    string `json:"grant_type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{}
	payload.GrantType = gt
	if len(args) < 2 {
		return nil, fmt.Errorf("Not enough arguments in call to 'Authenticate()'")
	} else if len(args) > 2 {
		return nil, fmt.Errorf("Too many arguments in call to 'Authenticate()'")
	}
	payload.ClientID = args[0]
	payload.ClientSecret = args[1]

	js, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}
	res, err := http.Post(Authentication(), "application/json", bytes.NewBuffer(js))

	if err != nil {
		return nil, err
	}
	if err = unmarshal(res, a); err != nil {

		return nil, err
	}
	return a, nil
}

func passwordGrant(gt string, args ...string) (*authentication, error) {
	a := &authentication{}

	payload := struct {
		GrantType    string `json:"grant_type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Username     string `json:"username"`
		Password     string `json:"password"`
	}{}
	payload.GrantType = gt
	if len(args) < 4 {
		return nil, fmt.Errorf("Not enough arguments in call to 'Authenticate()'")
	} else if len(args) > 4 {
		return nil, fmt.Errorf("Too many arguments in call to 'Authenticate()'")
	}
	payload.ClientID = args[0]
	payload.ClientSecret = args[1]
	payload.Username = args[2]
	payload.Password = args[3]

	js, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(Authentication(), "json", bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}
	if err = unmarshal(res, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (g *Cat) GetToken() string {
	return g.auth.Token
}
