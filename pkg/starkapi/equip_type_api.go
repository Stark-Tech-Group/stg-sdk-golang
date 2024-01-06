package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type EquipTypeApi struct {
	client *Client
}

func (api *EquipTypeApi) host() string {
	return api.client.host
}

func (api *EquipTypeApi) GetAll() (domain.EquipTypes, error) {
	url := fmt.Sprintf("%s/lists/equipType", api.host())

	var equipTypes domain.EquipTypes

	resp, err := api.client.get(url)
	if err != nil {
		return equipTypes, err
	}

	err = json.Unmarshal(resp, &equipTypes)
	if err != nil {
		return equipTypes, err
	}

	return equipTypes, nil
}
