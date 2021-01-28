package client

import (
	"bytes"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"io/ioutil"
	"net/http"
)

type Client struct{

	apiStatusEndpoint apiStatusEndpoint
	loginEndpoint loginEndpoint
	searchEndpoint searchEndpoint
	pointEndpoint pointEndpoint
	equipEndpoint EquipEndpoint
	siteEndpoint SiteEndpoint
	auth *response.AuthResponse
	httpClient *http.Client
	host string
}


func(client *Client) Init(host string) *Client{
	client.apiStatusEndpoint = apiStatusEndpoint{client}
	client.loginEndpoint = loginEndpoint{client}
	client.searchEndpoint = searchEndpoint{client}
	client.pointEndpoint = pointEndpoint{client}
	client.equipEndpoint = EquipEndpoint{client}
	client.siteEndpoint = SiteEndpoint{client}
	client.httpClient = &http.Client{}
	client.host = host

	fmt.Printf("Host: %s\n", client.host)

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
	status, err := client.apiStatusEndpoint.get()

	if err != nil {return nil, err }

	return status, nil
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

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	return client.doRequest(req)
}

func (client *Client) delete(url string) ([]byte, error){

	req, _ := http.NewRequest("DELETE", url, nil)

	client.setHeader(req)

	return client.doRequest(req)
}

func (client *Client) authPost(url string, requestBody []byte) ([]byte, error){
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	client.setHeader(req)

	return client.doRequest(req)

}

func (client *Client) setHeader(req *http.Request){
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s",client.auth.AccessToken))
	req.Header.Set("Content-Type", "application/json")
}

func(client *Client) doRequest( req *http.Request) ([]byte, error){
	resp, err := client.httpClient.Do(req)

	if err != nil{
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
