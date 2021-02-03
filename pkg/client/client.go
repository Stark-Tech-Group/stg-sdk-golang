package client

import (
	"bytes"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"io/ioutil"
	"net/http"
)

type Client struct{

	StatusApi     	StatusApi
	SearchApi     	SearchApi
	PointApi     	PointApi
	EquipApi      	EquipApi
	SiteApi       	SiteApi
	ProfileApi		ProfileApi
	/**/
	loginEndpoint 	authApi
	auth          *response.AuthResponse
	httpClient    *http.Client
	host          string
}


func(client *Client) Init(host string) *Client{
	client.loginEndpoint = authApi{client}
	client.httpClient = &http.Client{}
	client.host = host

	client.StatusApi 	= StatusApi{client:client}
	client.ProfileApi 	= ProfileApi{client: client}
	client.SearchApi 	= SearchApi{client:client}
	client.PointApi 	= PointApi{client:client}
	client.EquipApi 	= EquipApi{client:client}
	client.SiteApi 		= SiteApi{client:client}

	return client
}


func (client *Client) Login(un string, pw string) (*response.AuthResponse, error){
	login, err := client.loginEndpoint.login(un, pw)

	if err != nil{
		return nil, err
	}

	client.auth = login
	return client.auth, nil
}

func (client *Client) ApiStatus() (*response.StatusResponse, error) {
	status, err := client.StatusApi.get()

	if err != nil {return nil, err }

	return status, nil
}

func(client *Client) Search(body Query) (*response.SearchResponse, error) {
	search, err := client.SearchApi.Search(body)

	if err != nil{
		return nil, err
	}

	return search, nil
}

func(client *Client) DeletePoint(id int) *response.DeleteResponse {
	deleteResp, err := client.PointApi.delete(id)

	if err != nil {
		panic(err)
	}

	return deleteResp
}

func(client *Client) DeleteEquip(id int) *response.DeleteResponse {
	deleteResp, err := client.EquipApi.delete(id)

	if err != nil {
		panic(err)
	}

	return deleteResp
}

func(client *Client) DeleteSite(id int) *response.DeleteResponse {
	deleteResp, err := client.SiteApi.delete(id)

	if err != nil {
		panic(err)
	}

	return deleteResp
}



func (client *Client) get(url string) ([]byte, error){
	req, _ := http.NewRequest("GET", url, nil)
	client.setHeader(req)

	resp, err := client.httpClient.Do(req)

	if err != nil{
		return nil, err
	}

	if resp.StatusCode >= 299 {
		return nil, fmt.Errorf("unexpected response code [%s]", resp.Status)
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
