package starkapi

import (
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"net/http"
)

type MockClient struct {
	InitFunc        func(host string) *Client
	AuthFunc        func(accessToken, username string) (*response.AuthResponse, error)
	LoginFunc       func(un string, pw string) (*response.AuthResponse, error)
	ApiStatusFunc   func() (*response.StatusResponse, error)
	SearchFunc      func(body Query) (*response.SearchResponse, error)
	DeletePointFunc func(id uint32) (*response.DeleteResponse, error)
	DeleteEquipFunc func(id uint32) (*response.DeleteResponse, error)
	DeleteSiteFunc  func(id uint32) (*response.DeleteResponse, error)
	getFunc         func(url string) ([]byte, error)
	postFunc        func(url string, body []byte) ([]byte, error)
	putFunc         func(url string, requestBody []byte) ([]byte, error)
	deleteFunc      func(url string) ([]byte, error)
	setHeaderFunc   func(req *http.Request)
	doRequestFunc   func(req *http.Request) ([]byte, error)
	getHostFunc     func() string
}

func (client *MockClient) Init(host string) *Client {
	return client.InitFunc(host)
}

func (client *MockClient) Auth(accessToken, username string) (*response.AuthResponse, error) {
	return client.AuthFunc(accessToken, username)
}

func (client *MockClient) Login(un string, pw string) (*response.AuthResponse, error) {
	return client.LoginFunc(un, pw)
}

func (client *MockClient) ApiStatus() (*response.StatusResponse, error) {
	return client.ApiStatusFunc()
}

func (client *MockClient) Search(body Query) (*response.SearchResponse, error) {
	return client.SearchFunc(body)
}

func (client *MockClient) DeletePoint(id uint32) (*response.DeleteResponse, error) {
	return client.DeletePointFunc(id)
}

func (client *MockClient) DeleteEquip(id uint32) (*response.DeleteResponse, error) {
	return client.DeleteEquipFunc(id)
}

func (client *MockClient) DeleteSite(id uint32) (*response.DeleteResponse, error) {
	return client.DeleteSiteFunc(id)
}

func (client *MockClient) get(url string) ([]byte, error) {
	return client.getFunc(url)
}

func (client *MockClient) post(url string, body []byte) ([]byte, error) {
	return client.postFunc(url, body)
}

func (client *MockClient) put(url string, requestBody []byte) ([]byte, error) {
	return client.putFunc(url, requestBody)
}

func (client *MockClient) delete(url string) ([]byte, error) {
	return client.deleteFunc(url)
}

func (client *MockClient) setHeader(req *http.Request) {
	client.setHeaderFunc(req)
}

func (client *MockClient) doRequest(req *http.Request) ([]byte, error) {
	return client.doRequestFunc(req)
}

func (client *MockClient) getHost() string {
	return client.getHostFunc()
}
