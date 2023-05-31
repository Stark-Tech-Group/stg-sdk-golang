package domain

// Reports is a struct that represents a list of reports
type Reports struct {
	Count   int32     `json:"count"`
	Reports []*Report `json:"reports"`
}

// Report is a struct that represents a report
type Report struct {
	Id          int32  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Ref         string `json:"ref,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Description string `json:"description,omitempty"`
	Audit       *Audit `json:"audit,omitempty"`
}
