package client

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type SiteEndpoint struct{
	client *Client
}

func (siteEndpoint *SiteEndpoint) delete(id int) (*response.DeleteResponse, error){
	resp, err := siteEndpoint.client.delete(siteUrl(siteEndpoint.client.host, id))
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

func (siteEndpoint *SiteEndpoint) host() string {
	return siteEndpoint.client.host
}

func (siteEndpoint *SiteEndpoint) BaseUrl() string {
	return fmt.Sprintf("%s/core/sites", siteEndpoint.client.host)
}

func (siteEndpoint *SiteEndpoint) GetOne(id int) (domain.Site, error) {
	var site domain.Site

	resp, err := siteEndpoint.client.get(siteUrl(siteEndpoint.client.host, id))

	if err != nil { return site, err }

	err = json.Unmarshal(resp, &site)
	if err != nil{
		return site, err
	}

	return site, nil
}

func (siteEndpoint *SiteEndpoint) GetAll() (domain.Sites, error) {
	var sites domain.Sites
	url := siteEndpoint.BaseUrl() + "/"
	resp, err := siteEndpoint.client.get(url)

	if err != nil { return sites, err }




	err = json.Unmarshal(resp, &sites)

	fmt.Printf("Count: %s\n", sites.Count)

	if err != nil{
		return sites, err
	}

	return sites, nil
}
