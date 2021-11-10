package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type EquipApi struct {
	client *Client
}

func (equipApi *EquipApi) host() string {
	return equipApi.client.host
}

func (equipApi *EquipApi) delete(id uint32) (*response.DeleteResponse, error) {

	resp, err := equipApi.client.delete(equipUrl(equipApi.host(), id))
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

func equipsUrl(host string) string {
	return fmt.Sprintf("%s/core/equips", host)
}

func equipUrl(host string, id uint32) string {
	return fmt.Sprintf("%s/%d", equipsUrl(host), id)
}

func (equipApi *EquipApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/equips", equipApi.client.host)
}

func (equipApi *EquipApi) GetOne(id uint32) (domain.Equip, error) {
	url := fmt.Sprintf("%s/%v", equipApi.BaseUrl(), id)

	var equip domain.Equip

	resp, err := equipApi.client.get(url)
	if err != nil {
		return equip, err
	}
	err = json.Unmarshal(resp, &equip)
	if err != nil {
		return equip, err
	}

	return equip, nil
}

func (equipApi *EquipApi) GetAll() (domain.Equips, error) {
	url := fmt.Sprintf("%s/", equipApi.BaseUrl())

	var equips domain.Equips

	resp, err := equipApi.client.get(url)
	if err != nil {
		return equips, err
	}

	err = json.Unmarshal(resp, &equips)
	if err != nil {
		return equips, err
	}

	return equips, nil
}

func (equipApi *EquipApi) UpdateOne(id uint32, jsonBody []byte) (domain.Equip, error) {
	url := fmt.Sprintf("%s/%v", equipApi.BaseUrl(), id)

	var equip domain.Equip
	resp, err := equipApi.client.put(url, jsonBody)
	if err != nil {
		return equip, err
	}

	err = json.Unmarshal(resp, &equip)
	if err != nil {
		return equip, err
	}

	return equip, nil
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

func (equipApi *EquipApi) AddNewTag(equip *domain.Equip, name string, value string) error {
	url := fmt.Sprintf("%s/%v/tags", equipApi.BaseUrl(), equip.Id)
	ask := domain.Tag{Name: name, Value: value}

	body, err := json.Marshal(ask)
	if err != nil {
		return err
	}

	_, err = equipApi.client.post(url, body)
	if err != nil {
		return err
	}

	return nil
}

func (equipApi *EquipApi) CurVals(equipId uint32, urids []string) (domain.EquipCurVals, error) {
	url := fmt.Sprintf("%s/%v/curVals", equipApi.BaseUrl(), equipId)

	var equipCurVals domain.EquipCurVals

	data := struct {
		Urids []string `json:"urids"`
	}{urids}

	body, err := json.Marshal(data)
	if err != nil {
		return equipCurVals, err
	}

	resp, err := equipApi.client.post(url, body)
	if err != nil {
		return equipCurVals, err
	}

	err = json.Unmarshal(resp, &equipCurVals)
	if err != nil {
		return equipCurVals, err
	}

	return equipCurVals, nil
}
