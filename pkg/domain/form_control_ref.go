package domain

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	errNoControlNameProvided = "please provide a form control name"
	errNoRefProvided         = "please provide a ref"
	errNoValueProvided       = "please provide a value"
	errInvalidFormControl    = "invalid form control provided with name [%s]"
	errInvalidFormControlRef = "unable to validate form control ref with information provided"
)

type FormControlRefList struct {
	Count              int32             `json:"count"`
	FormControlRefList []*FormControlRef `json:"formControlList"`
}

type FormControlRef struct {
	Id            int32       `json:"id,omitempty"`
	TargetRef     string      `json:"targetRef" validate:"required"`
	SortOrder     int32       `json:"sortOrder" validate:"required"`
	Description   string      `json:"description" validate:"required"`
	Name          string      `json:"name" validate:"required"`
	Key           string      `json:"key" validate:"required"`
	Value         string      `json:"value" validate:"required"`
	FormControlId string      `json:"formControlId" validate:"required"`
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

func (o *FormControlRef) ValidateCreateRequireFields(formControlName string, ref string, value string) error {
	if formControlName == "" {
		return errors.New(errNoControlNameProvided)
	}
	if ref == "" {
		return errors.New(errNoRefProvided)
	}
	if value == "" {
		return errors.New(errNoValueProvided)
	}

	return nil
}

func (o *FormControlRef) BuildFormControlRefForCreate(formControl FormControl, ref string, value string) error {
	err := formControl.Validate()
	if err != nil {
		return errors.New(fmt.Sprintf(errInvalidFormControl, formControl.Name))
	}
	err = json.Unmarshal([]byte(formControl.Control), &formControl.ControlJSON)
	if err != nil {
		return errors.New(fmt.Sprintf(errInvalidFormControl, formControl.Name))
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
		return errors.New(errInvalidFormControlRef)
	}
	return nil
}
