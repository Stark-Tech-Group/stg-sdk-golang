package domain

type PasswordValidation struct {
	validated	bool `json:"validated"`
	err     	[]error `json:"errors"`
}

