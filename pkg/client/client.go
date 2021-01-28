package client

import (
	"bytes"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/env"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct{

	apiStatusEndpoint apiStatusEndpoint
	loginEndpoint     loginEndpoint
	searchEndpoint    searchEndpoint
	pointEndpoint     pointEndpoint
	equipEndpoint     EquipEndpoint
	siteEndpoint      SiteEndpoint
	auth              *response.AuthResponse
	httpClient        *http.Client
	host              string
}

func(client *Client) Init() *Client{
	client.apiStatusEndpoint = apiStatusEndpoint{client}
	client.loginEndpoint = loginEndpoint{client}
	client.searchEndpoint = searchEndpoint{client}
	client.pointEndpoint = pointEndpoint{client}
	client.equipEndpoint = EquipEndpoint{client}
	client.equipEndpoint = EquipEndpoint{client}
	client.siteEndpoint = SiteEndpoint{client}
	client.httpClient = &http.Client{}
	client.host = os.Getenv(env.STG_SDK_API_HOST)

	fmt.Printf("Host: %s\n", client.host)

	return client
}


func (client *Client) Login(un string, pw string){
	login, err := client.loginEndpoint.login(un, pw)

	if err != nil{
		panic(err)
	}

	client.auth = login
}

func (client *Client) ApiStatus() *response.StatusResponse {
	status, err := client.apiStatusEndpoint.get()

	if err != nil{
		panic(err)
	}

	return status
}

func(client *Client) Search(body SearchBody) *response.SearchResponse {
	search, err := client.searchEndpoint.search(body)

	if err != nil{
		panic(err)
	}

	return search
}

func(client *Client) DeletePoint(id int) *response.DeleteResponse {
	deleteResp, err := client.pointEndpoint.delete(id)

	if err != nil {
		panic(err)
	}

	return deleteResp
}

func(client *Client) DeleteEquip(id int) *response.DeleteResponse {
	deleteResp, err := client.equipEndpoint.delete(id)

	if err != nil {
		panic(err)
	}

	return deleteResp
}

func(client *Client) DeleteSite(id int) *response.DeleteResponse {
	deleteResp, err := client.siteEndpoint.delete(id)

	if err != nil {
		panic(err)
	}

	return deleteResp
}



func (client *Client) get(url string) ([]byte, error){
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.httpClient.Do(req)

	if err != nil{
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (client *Client) post(url string, requestBody []byte) ([]byte, error){
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	if err != nil{
		return nil, err
	}

	return client.doRequest(req)
}

func (client *Client) delete(url string) ([]byte, error){
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil{
		return nil, err
	}

	return client.doRequest(req)
}

func (client *Client) authPost(url string, requestBody []byte) ([]byte, error){
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil{
		return nil, err
	}

	return client.doRequest(req)
}

func (client *Client) setHeader(req *http.Request){
	if client.auth != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.auth.AccessToken))
		req.Header.Set("Content-Type", "application/json")
	}
}

func(client *Client) doRequest( req *http.Request) ([]byte, error){
	client.setHeader(req)
	resp, err := client.httpClient.Do(req)

	if err != nil{
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func(client *Client) GetSiteEndpoint() SiteEndpoint {
	return client.siteEndpoint
}
