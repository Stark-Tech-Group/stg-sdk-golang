package client

import (
	"encoding/json"
	"fmt"
	"starktechgroup/stg-sdk-golang/pkg/api/response"
)

type loginEndpoint struct{
	client *Client
}

func (loginEndpoint *loginEndpoint) login(un string, pw string) (*response.AuthResponse, error){
	requestBody, err := json.Marshal(map[string]string{
		"username": un,
		"password": pw,
	})

	if err != nil{
		return nil, err
	}

	resp, err := loginEndpoint.client.post(loginUrl(loginEndpoint.client.host), requestBody)

	login := response.AuthResponse{}
	err = json.Unmarshal(resp, &login)

	if err != nil{
		panic(err)
	}

	return &login, nil
}

func loginUrl(host string) string{
	return fmt.Sprintf("%s/login", host)
}
