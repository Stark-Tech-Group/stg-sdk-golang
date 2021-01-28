package client

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
)

type apiStatusEndpoint struct {
	client *Client
}

func (endpoint *apiStatusEndpoint) get() (*response.StatusResponse, error){
	resp, err := endpoint.client.get(apiStatusUrl(endpoint.client.host))

	if err != nil{
		return nil, err
	}

	apiStatus := response.StatusResponse{}
	err = json.Unmarshal(resp, &apiStatus)

	if err != nil{
		panic(err)
	}

	return &apiStatus, nil
}

func apiStatusUrl(host string) string{
	return fmt.Sprintf("%s/apistatus", host)
}
