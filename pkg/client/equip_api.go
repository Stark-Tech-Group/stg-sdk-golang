package client

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type EquipApi struct{
	client *Client
}

func (equipApi *EquipApi) host() string {
	return equipApi.client.host
}

func (equipApi *EquipApi) delete(id int) (*response.DeleteResponse, error){

	resp, err := equipApi.client.delete(equipUrl(equipApi.host(), id))
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

func equipsUrl(host string) string {
	return fmt.Sprintf("%s/core/equips",host)
}

func equipUrl(host string , id int) string {
	return fmt.Sprintf("%s/%d", equipsUrl(host), id)
}

func (equipApi *EquipApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/equips", equipApi.client.host)
}

func (equipApi *EquipApi) CreateOne(ask domain.Equip) (domain.Equip, error) {
	url := equipApi.BaseUrl()
	body, err := json.Marshal(ask)

	var equip domain.Equip
	if err != nil {
		return equip, err
	}

	resp, err := equipApi.client.post(url, body)
	if err != nil {
		return equip, err
	}
	err = json.Unmarshal(resp, &equip)
	if err != nil {
		return equip, err
	}

	return equip, nil
}
