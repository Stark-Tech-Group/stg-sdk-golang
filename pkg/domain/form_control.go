package domain

import (
	"github.com/go-playground/validator/v10"
)

type FormControlList struct {
	Count           int32          `json:"count"`
	FormControlList []*FormControl `json:"formControlList"`
}

type FormControl struct {
	Id          int32  `json:"id,omitempty"`
	Ref         string `json:"ref,omitempty" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Enabled     bool   `json:"enabled,omitempty"`
	Description string `json:"description"`
	Control     string `json:"string" validate:"required"`
	Audit       *Audit `json:"audit,omitempty"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (o *FormControl) Validate() error {
	return validate.Struct(o)
}
