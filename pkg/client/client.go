package client

import (
	"bytes"
	"fmt"
	"go-scripts/internal/response"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct{

	apiStatusEndpoint apiStatusEndpoint
	loginEndpoint loginEndpoint
	searchEndpoint searchEndpoint
	pointEndpoint pointEndpoint
	equipEndpoint equipEndpoint
	siteEndpoint siteEndpoint
	auth *response.Auth
	httpClient *http.Client
	host string
}


func(client *Client) Init() *Client{
	client.apiStatusEndpoint = apiStatusEndpoint{client}
	client.loginEndpoint = loginEndpoint{client}
	client.searchEndpoint = searchEndpoint{client}
	client.pointEndpoint = pointEndpoint{client}
	client.equipEndpoint = equipEndpoint{client}
	client.equipEndpoint = equipEndpoint{client}
	client.siteEndpoint = siteEndpoint{client}
	client.httpClient = &http.Client{}
	client.host = os.Getenv("host")

	fmt.Printf("Host: %s\n", client.host)

	return client
}


func (client *Client) Login(un string, pw string){
	login, err := client.loginEndpoint.login(un, pw)

	if err != nil{
		panic(err)
	}

	fmt.Printf("Login: %s\n", login.AccessToken)
	client.auth = login
}

func (client *Client) ApiStatus() *response.ApiStatus{
	status, err := client.apiStatusEndpoint.get()

	if err != nil{
		panic(err)
	}

	return status
}

func(client *Client) Search(body SearchBody) *response.Search{
	search, err := client.searchEndpoint.search(body)

	if err != nil{
		panic(err)
	}

	return search
}

func(client *Client) DeletePoint(id int) *response.Delete{
	deleteResp, err := client.pointEndpoint.delete(id)

	if err != nil {
		panic(err)
	}

	return deleteResp
}

func(client *Client) DeleteEquip(id int) *response.Delete{
	deleteResp, err := client.equipEndpoint.delete(id)

	if err != nil {
		panic(err)
	}

	return deleteResp
}

func(client *Client) DeleteSite(id int) *response.Delete{
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
