package domain

type FormControlList struct {
	Count  int32    `json:"count"`
	FormControlList []*FormControl `json:"formControlList"`
}

type FormControl struct {
	Id          int32      `json:"id,omitempty"`
	Ref         string     `json:"ref,omitempty"`
	Name        string     `json:"name"`
	Enabled     bool       `json:"enabled,omitempty"`
	Description string     `json:"description"`
	Control 	string     `json:"string"`
	Audit       *Audit     `json:"audit,omitempty"`
}
