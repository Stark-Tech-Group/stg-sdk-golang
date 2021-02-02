package client

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
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
