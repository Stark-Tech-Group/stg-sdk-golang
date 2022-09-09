package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type FormsApi struct {
	client *Client
}

func (formsApi *FormsApi) host() string {
	return formsApi.client.host
}

func (formsApi *FormsApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/forms", formsApi.host())
}

func (formsApi *FormsApi) ControlsPrefix() string {
	return fmt.Sprintf("/controls")
}

func (formsApi *FormsApi) GetAllControls() (domain.FormControlList, error) {
	url := fmt.Sprintf("%s%s", formsApi.BaseUrl(), formsApi.ControlsPrefix())

	var controls domain.FormControlList

	resp, err := formsApi.client.get(url)
	if err != nil {
		return controls, err
	}

	err = json.Unmarshal(resp, &controls)
	if err != nil {
		return controls, err
	}

	return controls, nil
}

func (formsApi *FormsApi) GetAllControlsForAsset(ref string) (domain.FormControlRefList, error) {
	url := fmt.Sprintf("%s/%s%s", formsApi.BaseUrl(), ref, formsApi.ControlsPrefix())

	var assetControls domain.FormControlRefList

	resp, err := formsApi.client.get(url)
	if err != nil {
		return assetControls, err
	}

	err = json.Unmarshal(resp, &assetControls)
	if err != nil {
		return assetControls, err
	}

	return assetControls, nil
}