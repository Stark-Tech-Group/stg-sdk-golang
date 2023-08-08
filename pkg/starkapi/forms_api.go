package starkapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
	logger "github.com/sirupsen/logrus"
)

const (
	errNoControlNameProvided      = "please provide a form control name"
	errNoRefProvided              = "please provide a ref"
	errNoValueProvided            = "please provide a value"
	errInvalidControlNameProvided = "no form control found with name provided : [%s]"
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

func (formsApi *FormsApi) CreateControlOnRef(formControlName string, ref string, value string) (domain.FormControlRef, error) {
	var control domain.FormControl
	var controlRef domain.FormControlRef

	err := controlRef.ValidateStringParams(formControlName, errNoControlNameProvided)
	if err != nil {
		logger.Errorf("controlRef.ValidateStringParams failed with error : [%s]", err)
		return controlRef, err
	}

	err = controlRef.ValidateStringParams(ref, errNoRefProvided)
	if err != nil {
		logger.Errorf("controlRef.ValidateStringParams failed with error : [%s]", err)
		return controlRef, err
	}

	err = controlRef.ValidateStringParams(value, errNoValueProvided)
	if err != nil {
		logger.Errorf("controlRef.ValidateStringParams failed with error : [%s]", err)
		return controlRef, err
	}

	control, err = formsApi.GetControlByName(formControlName)
	if err != nil {
		logger.Errorf("formsApi.GetControlByName failed with error : [%s]", err)
		return controlRef, err
	}

	if control.Id < 1 || control.Name != formControlName {
		newError := errors.New(fmt.Sprintf(errInvalidControlNameProvided, formControlName))
		logger.Error(newError)
		return controlRef, newError
	}

	err = controlRef.BuildFormControlRefForCreate(control, ref, value)
	if err != nil {
		logger.Errorf("controlRef.BuildFormControlRefForCreate failed with error : [%s]", err)
		return controlRef, err
	}

	url := fmt.Sprintf("%s/%s%s", formsApi.baseUrl(), ref, formsApi.controlsPrefix())

	body, err := json.Marshal(controlRef)
	if err != nil {
		logger.Errorf("failed to unmarshall controlRef json with error : [%s]", err)
		return controlRef, err
	}

	resp, err := formsApi.client.post(url, body)
	if err != nil {
		logger.Errorf("failed to run formsApi.client.post with error : [%s]", err)
		return controlRef, err
	}

	if len(resp) > 0 {
		err = json.Unmarshal(resp, &controlRef)
		if err != nil {
			logger.Errorf("failed to unmarshall post response json with error : [%s]", err)
			return controlRef, err
		}
	}

	return controlRef, nil
}
