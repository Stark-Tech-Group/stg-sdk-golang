package client

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
)

type searchEndpoint struct{
	client *Client
}

type SearchBody struct {
	Query       string `json:"query"`
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"pageSize"`
}

func(searchEndpoint *searchEndpoint) search (searchBody SearchBody) (*response.SearchResponse, error){
	requestBody, err := json.Marshal(searchBody)

	if err != nil{
		return nil, err
	}

	resp, err := searchEndpoint.client.authPost(searchUrl(searchEndpoint.client.host), requestBody)

	if err != nil{
		return nil, err
	}

	search := response.SearchResponse{}

	err = json.Unmarshal(resp, &search)

	if err != nil{
		panic(err)
	}

	return &search, nil
}

func searchUrl(host string) string{
	return fmt.Sprintf("%s/core/search/assets", host)
}
