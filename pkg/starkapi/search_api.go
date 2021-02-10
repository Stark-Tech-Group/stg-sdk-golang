package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
)

type SearchApi struct{
	client *Client
}

type Query struct {
	Query       string `json:"query"`
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"pageSize"`
}

/*
convenience func to Search(Query)
 */
func(searchApi *SearchApi) SearchText(query string, page int, size int) (*response.SearchResponse, error){
	return searchApi.Search(Query{
		Query: query,
		 CurrentPage: page,
		 PageSize: size,
	})
}

/**
Search using a query
 */
func(searchApi *SearchApi) Search(query Query) (*response.SearchResponse, error){
	requestBody, err := json.Marshal(query)

	if err != nil{
		return nil, err
	}

	resp, err := searchApi.client.post(searchUrl(searchApi.client.host), requestBody)

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
