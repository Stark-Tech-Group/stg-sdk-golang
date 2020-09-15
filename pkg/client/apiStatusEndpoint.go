package client

import (
	"encoding/json"
	"fmt"
	"go-scripts/internal/response"
)

type apiStatusEndpoint struct {
	client *Client
}

func (endpoint *apiStatusEndpoint) get() (*response.ApiStatus, error){
	resp, err := endpoint.client.get(apiStatusUrl(endpoint.client.host))

	if err != nil{
		return nil, err
	}

	apiStatus := response.ApiStatus{}
	err = json.Unmarshal(resp, &apiStatus)

	if err != nil{
		panic(err)
	}

	return &apiStatus, nil
}

func apiStatusUrl(host string) string{
	return fmt.Sprintf("%s/apistatus", host)
}
