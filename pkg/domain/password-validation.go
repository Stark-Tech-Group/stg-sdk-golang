package domain

type PasswordValidation struct {
	Validated bool    `json:"validated"`
	Err       []error `json:"errors"`
}
