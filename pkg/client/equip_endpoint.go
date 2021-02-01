package client

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
)

type EquipEndpoint struct{
	client *Client
}

func (equipEndpoint *EquipEndpoint) host() string {
	return equipEndpoint.client.host
}

func (equipEndpoint *EquipEndpoint) delete(id int) (*response.DeleteResponse, error){

	resp, err := equipEndpoint.client.delete(equipUrl(equipEndpoint.host(), id))
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