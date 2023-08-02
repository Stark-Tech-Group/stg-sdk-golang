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

func (formsApi *FormsApi) baseUrl() string {
	return fmt.Sprintf("%s/core/forms", formsApi.host())
}

func (formsApi *FormsApi) controlsPrefix() string {
	return fmt.Sprintf("/controls")
}

func (formsApi *FormsApi) GetAllControls() (domain.FormControlList, error) {
	url := fmt.Sprintf("%s%s", formsApi.baseUrl(), formsApi.controlsPrefix())

	var controls domain.FormControlList

	resp, err := formsApi.client.get(url)
	if err != nil {
		return controls, err
	}

	if len(resp) > 0 {
		err = json.Unmarshal(resp, &controls)
		if err != nil {
			return controls, err
		}
	}

	return controls, nil
}

func (formsApi *FormsApi) GetAllControlsForAsset(ref string) (domain.FormControlRefList, error) {
	url := fmt.Sprintf("%s/%s%s", formsApi.baseUrl(), ref, formsApi.controlsPrefix())

	var assetControls domain.FormControlRefList

	resp, err := formsApi.client.get(url)
	if err != nil {
		return assetControls, err
	}

	if len(resp) > 0 {
		err = json.Unmarshal(resp, &assetControls)
		if err != nil {
			return assetControls, err
		}
	}

	return assetControls, nil
}

func (formsApi *FormsApi) GetControlByName(name string) (domain.FormControl, error) {
	var control domain.FormControl
	var controlsList, err = formsApi.GetAllControls()

	if err != nil {
		return control, err
	}

	for i := range controlsList.FormControlList {
		if controlsList.FormControlList[i].Name == name {
			control = *controlsList.FormControlList[i]
			break
		}
	}

	return control, nil
}
