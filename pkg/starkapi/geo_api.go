package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type GeoApi struct {
	client *Client
}

func (geoApi *GeoApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/geo", geoApi.client.host)
}

func (geoApi *GeoApi) host() string {
	return geoApi.client.host
}

/*
returns geocoding for an address
*/
func (geoApi *GeoApi) GeoCoding(address string, city string, state string, postalCode string) (domain.GeoCoding, error) {
	var geoCoding domain.GeoCoding
	url := geoApi.BaseUrl()

	data := struct {
		Address    string `json:"county"`
		City       string `json:"city"`
		State      string `json:"state"`
		PostalCode string `json:"postalCode"`
	}{Address: address, City: city, State: state, PostalCode: postalCode}

	body, err := json.Marshal(data)
	if err != nil {
		return geoCoding, err
	}

	resp, err := geoApi.client.post(url, body)

	if err != nil {
		return geoCoding, err
	}

	err = json.Unmarshal(resp, &geoCoding)

	if err != nil {
		return geoCoding, err
	}

	return geoCoding, nil
}

/*
returns climate zone for a county and state
*/
func (geoApi *GeoApi) ClimateZone(county string, state string) (domain.ClimateZone, error) {
	var climateZone domain.ClimateZone
	url := fmt.Sprintf("%s/climateZone", geoApi.BaseUrl())

	data := struct {
		County string `json:"county"`
		State  string `json:"state"`
	}{County: county, State: state}

	body, err := json.Marshal(data)
	if err != nil {
		return climateZone, err
	}

	resp, err := geoApi.client.post(url, body)

	if err != nil {
		return climateZone, err
	}

	err = json.Unmarshal(resp, &climateZone)

	if err != nil {
		return climateZone, err
	}

	return climateZone, nil
}
