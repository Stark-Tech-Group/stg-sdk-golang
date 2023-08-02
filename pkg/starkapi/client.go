package starkapi

import (
	"bytes"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AssetTreeApi AssetTreeApi
	AssetsApi    AssetsApi
	StatusApi    StatusApi
	SearchApi    SearchApi
	PointApi     PointApi
	EquipApi     EquipApi
	SiteApi      SiteApi
	ProfileApi   ProfileApi
	ConnApi      ConnApi
	GeoApi       GeoApi
	UridApi      UridApi
	TagApi       TagApi
	FormsApi     FormsApi
	/**/
	loginEndpoint authApi
	auth          *response.AuthResponse
	httpClient    *http.Client
	host          string
}

func (client *Client) Init(host string) *Client {
	client.loginEndpoint = authApi{client}
	client.httpClient = &http.Client{}
	client.host = host

	client.StatusApi = StatusApi{client: client}
	client.ProfileApi = ProfileApi{client: client}
	client.SearchApi = SearchApi{client: client}
	client.PointApi = PointApi{client: client}
	client.EquipApi = EquipApi{client: client}
	client.SiteApi = SiteApi{client: client}
	client.ConnApi = ConnApi{client: client}
	client.AssetTreeApi = AssetTreeApi{client: client}
	client.AssetsApi = AssetsApi{client: client}
	client.GeoApi = GeoApi{client: client}
	client.UridApi = UridApi{client: client}
	client.TagApi = TagApi{client: client}
	client.FormsApi = FormsApi{client: client}

	return client
}

// Auth
func (client *Client) Auth(accessToken, username string) (*response.AuthResponse, error) {

	r := response.AuthResponse{
		AccessToken: accessToken,
		Username:    username,
	}

	client.auth = &r
	return client.auth, nil
}

func (client *Client) Login(un string, pw string) (*response.AuthResponse, error) {
	login, err := client.loginEndpoint.login(un, pw)

	if err != nil {
		return nil, err
	}

	client.auth = login
	return client.auth, nil
}

func (client *Client) ApiStatus() (*response.StatusResponse, error) {
	status, err := client.StatusApi.Get()

	if err != nil {
		return nil, err
	}

	return status, nil
}

func (client *Client) Search(body Query) (*response.SearchResponse, error) {
	search, err := client.SearchApi.Search(body)

	if err != nil {
		return nil, err
	}

	return search, nil
}

func (client *Client) DeletePoint(id uint32) (*response.DeleteResponse, error) {
	deleteResp, err := client.PointApi.delete(id)

	if err != nil {
		return nil, err
	}

	return deleteResp, nil
}

func (client *Client) DeleteEquip(id uint32) (*response.DeleteResponse, error) {
	deleteResp, err := client.EquipApi.delete(id)

	if err != nil {
		return nil, err
	}

	return deleteResp, nil
}

func (client *Client) DeleteSite(id uint32) (*response.DeleteResponse, error) {
	deleteResp, err := client.SiteApi.delete(id)

	if err != nil {
		return nil, err
	}

	return deleteResp, nil
}

func (client *Client) get(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	client.setHeader(req)

	resp, err := client.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 299 {
		return nil, fmt.Errorf("unexpected response code [%s]", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func (client *Client) post(url string, requestBody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	return client.doRequest(req)
}

func (client *Client) put(url string, requestBody []byte) ([]byte, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	return client.doRequest(req)
}

func (client *Client) delete(url string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return client.doRequest(req)
}

func (client *Client) setHeader(req *http.Request) {
	if client.auth != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.auth.AccessToken))
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Client-ID", "sdk-golang")
}

func (client *Client) doRequest(req *http.Request) ([]byte, error) {
	client.setHeader(req)
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 0 && resp.StatusCode < 300 {
		// ok status
		return ioutil.ReadAll(resp.Body)
	}

	return nil, fmt.Errorf("unexpected response code [%s]", resp.Status)
}

func (client *Client) getHost() string {
	return client.host
}
