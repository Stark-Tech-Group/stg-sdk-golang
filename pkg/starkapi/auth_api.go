package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
)

const (
	username = "username"
	password = "password"
)

type authApi struct {
	client *Client
}

func (authApi *authApi) login(un string, pw string) (*response.AuthResponse, error) {
	requestBody, err := json.Marshal(map[string]string{
		username: un,
		password: pw,
	})

	if err != nil {
		return nil, err
	}

	resp, err := authApi.client.post(loginUrl(authApi.client.host), requestBody)

	login := response.AuthResponse{}
	err = json.Unmarshal(resp, &login)

	if err != nil {
		return nil, err
	}

	return &login, nil
}

func loginUrl(host string) string {
	return fmt.Sprintf("%s/login", host)
}

func meUrl(host string) string {
	return fmt.Sprintf("%s/core/persons/me", host)
}

//RefreshKeychain refreshes the user's keychain
func (authApi *authApi) RefreshKeychain() error {
	url := meUrl(authApi.client.host) + "/refreshKeychain"

	_, err := authApi.client.post(url, nil)

	if err != nil {
		return err
	}

	return nil
}
