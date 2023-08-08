package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	logger "github.com/sirupsen/logrus"
)

const (
	errInvalidFormControl = "form control struct validate failed for form control name [%s], with error : [%s]"
)

type FormControlRefList struct {
	Count              int32             `json:"count"`
	FormControlRefList []*FormControlRef `json:"formControlList"`
}

type FormControlRef struct {
	Id            int64       `json:"id,omitempty"`
	TargetRef     string      `json:"targetRef" validate:"required"`
	SortOrder     int32       `json:"sortOrder" validate:"required"`
	Description   string      `json:"description" validate:"required"`
	Name          string      `json:"name" validate:"required"`
	Key           string      `json:"key" validate:"required"`
	Value         string      `json:"value" validate:"required"`
	FormControlId int64       `json:"formControlId" validate:"required"`
	FormControl   FormControl `json:"formControl,omitempty"`
	Audit         *Audit      `json:"audit,omitempty"`
	Ref           string      `json:"ref,omitempty"`
	Version       int32       `json:"version,omitempty"`
}

func NewFormControlRef() *FormControlRef {
	return &FormControlRef{
		Ref: NewRef(FormControlRefTable),
	}
}

func (o *FormControlRef) ValidateStringParams(paramName string, errString string) error {
	if paramName == "" {
		logger.Error(errors.New(errString))
		return errors.New(errString)
	}

	return nil
}

func (o *FormControlRef) BuildFormControlRefForCreate(formControl FormControl, ref string, value string) error {
	err := formControl.Validate()
	if err != nil {
		validateErr := errors.New(fmt.Sprintf(errInvalidFormControl, formControl.Name, err))
		logger.Error(validateErr)
		return validateErr
	}
	err = json.Unmarshal([]byte(formControl.Control), &formControl.ControlJSON)
	if err != nil {
		logger.Errorf("failed to unmarshall formControl.Control json with error : [%s]", err)
		return err
	}
	o.TargetRef = ref
	o.SortOrder = 1
	o.Description = formControl.Description
	o.Name = formControl.Name
	o.Key = formControl.ControlJSON.Key
	o.Value = value
	o.FormControl = formControl
	o.FormControlId = formControl.Id

	err = validate.Struct(o)
	if err != nil {
		logger.Errorf("form control ref struct validate failed with error : [%s]", err)
		return err
	}
	return nil
}
