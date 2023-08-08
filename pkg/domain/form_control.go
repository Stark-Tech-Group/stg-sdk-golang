package domain

import (
	"github.com/go-playground/validator/v10"
)

type FormControlList struct {
	Count           int32          `json:"count"`
	FormControlList []*FormControl `json:"formControlList"`
}

type FormControl struct {
	Id          int64  `json:"id,omitempty"`
	Ref         string `json:"ref" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Enabled     bool   `json:"enabled,omitempty"`
	Description string `json:"description,omitempty"`
	Control     string `json:"control" validate:"required"`
	ControlJSON struct {
		Key             string `json:"key"`
		Type            string `json:"type"`
		TemplateOptions struct {
			Label       string `json:"label"`
			Placeholder string `json:"placeholder"`
			Required    bool   `json:"required"`
		} `json:"templateOptions,omitempty"`
	} `json:"-,omitempty"`
	Version string `json:"version,omitempty"`
	Audit   *Audit `json:"audit,omitempty"`
	Links   struct {
		Type string `json:"type"`
		Self string `json:"self"`
	} `json:"links,omitempty"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (o *FormControl) Validate() error {
	return validate.Struct(o)
}
