package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type UridApi struct {
	client *Client
}

func (uridApi *UridApi) host() string {
	return uridApi.client.host
}

func (uridApi *UridApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/pointUrids", uridApi.client.host)
}

func (uridApi *UridApi) AdminUrl() string {
	return fmt.Sprintf("%s/admin/pointUrids", uridApi.client.host)
}

func (uridApi *UridApi) GetOne(id uint32) (domain.Urid, error) {
	var urid domain.Urid

	resp, err := uridApi.client.get(fmt.Sprintf("%s/%d", uridApi.BaseUrl(), id))

	if err != nil {
		return urid, err
	}

	err = json.Unmarshal(resp, &urid)
	if err != nil {
		return urid, err
	}

	return urid, nil
}

func (uridApi *UridApi) GetAll() (domain.Urids, error) {
	var urids domain.Urids
	url := uridApi.BaseUrl() + "/"
	resp, err := uridApi.client.get(url)

	if err != nil {
		return urids, err
	}

	err = json.Unmarshal(resp, &urids)

	if err != nil {
		return urids, err
	}

	return urids, nil
}

func (uridApi *UridApi) CreateOne(ask domain.Urid) (domain.Urid, error) {
	url := uridApi.AdminUrl()
	body, err := json.Marshal(ask)

	var urid domain.Urid
	if err != nil {
		return urid, err
	}

	resp, err := uridApi.client.post(url, body)
	if err != nil {
		return urid, err
	}
	err = json.Unmarshal(resp, &urid)
	if err != nil {
		return urid, err
	}

	return urid, nil
}

func (uridApi *UridApi) UpdateOne(id uint32, jsonBody []byte) (domain.Urid, error) {
	url := fmt.Sprintf("%s/%v", uridApi.AdminUrl(), id)

	var urid domain.Urid
	resp, err := uridApi.client.put(url, jsonBody)
	if err != nil {
		return urid, err
	}

	err = json.Unmarshal(resp, &urid)
	if err != nil {
		return urid, err
	}

	return urid, nil
}
