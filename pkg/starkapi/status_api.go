package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
)

type StatusApi struct {
	client *Client
}

func (statusApi *StatusApi) Get() (*response.StatusResponse, error){
	resp, err := statusApi.client.get(apiStatusUrl(statusApi.client.host))

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
