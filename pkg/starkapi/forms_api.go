package starkapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

const (
	errNoRefProvided = "a valid ref must be provided for this call"
)

type FormsApi struct {
	client ApiClient
}

func (formsApi *FormsApi) host() string {
	return formsApi.client.getHost()
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

func (formsApi *FormsApi) CreateControlOnRef(control domain.FormControl) (domain.FormControl, error) {
	if control.Ref == "" {
		return control, errors.New(errNoRefProvided)
	}
	url := fmt.Sprintf("%s/%s%s", formsApi.baseUrl(), control.Ref, formsApi.controlsPrefix())

	body, err := json.Marshal(control)
	if err != nil {
		return control, err
	}

	resp, err := formsApi.client.post(url, body)
	if err != nil {
		return control, err
	}
	err = json.Unmarshal(resp, &control)
	if err != nil {
		return control, err
	}

	return control, nil
}
