package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type SiteApi struct{
	client *Client
}

func (siteApi *SiteApi) delete(id int) (*response.DeleteResponse, error){
	resp, err := siteApi.client.delete(siteUrl(siteApi.client.host, id))
	if err != nil{
		return nil, err
	}

	deleteResp := response.DeleteResponse{}
	err = json.Unmarshal(resp, &deleteResp)

	if err != nil{
		panic(err)
	}

	return &deleteResp, nil
}

func sitesUrl(host string) string {
	return fmt.Sprintf("%s/core/sites",host)
}

func siteUrl(host string , id int) string {
	return fmt.Sprintf("%s/%d", sitesUrl(host), id)
}

func (siteApi *SiteApi) host() string {
	return siteApi.client.host
}

func (siteApi *SiteApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/sites", siteApi.client.host)
}

/*
returns one site provided the id
 */
func (siteApi *SiteApi) GetOne(id int) (domain.Site, error) {
	var site domain.Site

	resp, err := siteApi.client.get(siteUrl(siteApi.client.host, id))

	if err != nil { return site, err }

	err = json.Unmarshal(resp, &site)
	if err != nil{
		return site, err
	}

	return site, nil
}

/*
returns all the sites the current auth has access to
 */
func (siteApi *SiteApi) GetAll() (domain.Sites, error) {
	var sites domain.Sites
	url := siteApi.BaseUrl() + "/"
	resp, err := siteApi.client.get(url)

	if err != nil { return sites, err }

	err = json.Unmarshal(resp, &sites)

	if err != nil{
		return sites, err
	}

	return sites, nil
}

func (siteApi *SiteApi) CreateOne(ask domain.Site) (domain.Site, error) {
	url := siteApi.BaseUrl()
	body, err := json.Marshal(ask)

	var site domain.Site
	if err != nil {
		return site, err
	}

	resp, err := siteApi.client.post(url, body)
	if err != nil {
		return site, err
	}
	err = json.Unmarshal(resp, &site)
	if err != nil {
		return site, err
	}

	return site, nil
}

func (siteApi *SiteApi) UpdateOne(ask domain.Site) (domain.Site, error) {
	url := fmt.Sprintf("%s/%v", siteApi.BaseUrl(), ask.Id)
	body, err := json.Marshal(ask)

	var site domain.Site
	if err != nil { return site, err }

	resp, err := siteApi.client.put(url, body)
	if err != nil { return site, err }

	err = json.Unmarshal(resp, &site)
	if err != nil { return site, err }

	return site, nil
}

func (siteApi *SiteApi) AddNewTag(site domain.Site, name string, value string) error {
	url := fmt.Sprintf("%s/%v/tags", siteApi.BaseUrl(), site.Id)
	ask := domain.Tag{Name: name, Value: value}

	body, err := json.Marshal(ask)
	if err != nil {
		return err
	}

	_, err = siteApi.client.post(url, body)
	if err != nil {
		return err
	}

	return  nil
}
