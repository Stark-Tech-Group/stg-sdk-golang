package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type SiteApi struct {
	client *Client
}

func (api *SiteApi) delete(id uint32) (*response.DeleteResponse, error) {
	resp, err := api.client.delete(siteUrl(api.client.host, id))
	if err != nil {
		return nil, err
	}

	deleteResp := response.DeleteResponse{}
	err = json.Unmarshal(resp, &deleteResp)

	if err != nil {
		return nil, err
	}

	return &deleteResp, nil
}

func sitesUrl(host string) string {
	return fmt.Sprintf("%s/core/sites", host)
}

func siteUrl(host string, id uint32) string {
	return fmt.Sprintf("%s/%d", sitesUrl(host), id)
}

func (api *SiteApi) host() string {
	return api.client.host
}

func (api *SiteApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/sites", api.client.host)
}

/*
returns one site provided the id
*/
func (api *SiteApi) GetOne(id uint32) (domain.Site, error) {
	var site domain.Site

	resp, err := api.client.get(siteUrl(api.client.host, id))

	if err != nil {
		return site, err
	}

	err = json.Unmarshal(resp, &site)
	if err != nil {
		return site, err
	}

	return site, nil
}

/*
returns all the sites the current auth has access to
*/
func (api *SiteApi) GetAll() (domain.Sites, error) {
	var sites domain.Sites
	url := api.BaseUrl() + "/"
	resp, err := api.client.get(url)

	if err != nil {
		return sites, err
	}

	err = json.Unmarshal(resp, &sites)

	if err != nil {
		return sites, err
	}

	return sites, nil
}

func (api *SiteApi) CreateOne(ask domain.Site) (domain.Site, error) {
	url := api.BaseUrl()
	body, err := json.Marshal(ask)

	var site domain.Site
	if err != nil {
		return site, err
	}

	resp, err := api.client.post(url, body)
	if err != nil {
		return site, err
	}
	err = json.Unmarshal(resp, &site)
	if err != nil {
		return site, err
	}

	return site, nil
}

func (api *SiteApi) UpdateOne(id uint32, jsonBody []byte) (domain.Point, error) {
	url := fmt.Sprintf("%s/%v", api.BaseUrl(), id)

	var point domain.Point
	resp, err := api.client.put(url, jsonBody)
	if err != nil {
		return point, err
	}

	err = json.Unmarshal(resp, &point)
	if err != nil {
		return point, err
	}

	return point, nil
}

func (api *SiteApi) AddNewTag(site domain.Site, name string, value string) error {
	url := fmt.Sprintf("%s/%v/tags", api.BaseUrl(), site.Id)
	ask := domain.Tag{Name: name, Value: value}

	body, err := json.Marshal(ask)
	if err != nil {
		return err
	}

	_, err = api.client.post(url, body)
	if err != nil {
		return err
	}

	return nil
}

