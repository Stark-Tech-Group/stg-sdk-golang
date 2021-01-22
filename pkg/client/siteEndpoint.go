package client

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
)

type siteEndpoint struct{
	client *Client
}

func (siteEndpoint *siteEndpoint) delete(id int) (*response.DeleteResponse, error){

	resp, err := siteEndpoint.client.delete(siteUrl(siteEndpoint.client.host, id))

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

func sitesUrl(host string) string {
	return fmt.Sprintf("%s/core/sites",host)
}

func siteUrl(host string , id int) string {
	return fmt.Sprintf("%s/%d", sitesUrl(host), id)
}
