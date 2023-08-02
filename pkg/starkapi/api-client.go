package starkapi

import (
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"net/http"
)

type ApiClient interface {
	Init(host string) *Client
	Auth(accessToken, username string) (*response.AuthResponse, error)
	Login(un string, pw string) (*response.AuthResponse, error)
	ApiStatus() (*response.StatusResponse, error)
	Search(body Query) (*response.SearchResponse, error)
	DeletePoint(id uint32) (*response.DeleteResponse, error)
	DeleteEquip(id uint32) (*response.DeleteResponse, error)
	DeleteSite(id uint32) (*response.DeleteResponse, error)
	get(url string) ([]byte, error)
	post(url string, body []byte) ([]byte, error)
	put(url string, requestBody []byte) ([]byte, error)
	delete(url string) ([]byte, error)
	setHeader(req *http.Request)
	doRequest(req *http.Request) ([]byte, error)
	getHost() string
}
