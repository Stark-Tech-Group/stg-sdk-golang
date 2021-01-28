package client

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
)

type pointEndpoint struct{
	client *Client
}

func (pointEndpoint *pointEndpoint) delete(id int) (*response.DeleteResponse, error){

	resp, err := pointEndpoint.client.delete(pointUrl(pointEndpoint.client.host, id))

	if err != nil{
		return nil, err
	}

	deleteResp := response.DeleteResponse{}

	err = json.Unmarshal(resp, &deleteResp)

	if err != nil{
		panic(err)
	}

	return &deleteResp, nil
}

func pointsUrl(host string) string {
	return fmt.Sprintf("%s/core/points",host)
}

func pointUrl(host string , id int) string {
	return fmt.Sprintf("%s/%d", pointsUrl(host), id)
}
