package domain

type Audit struct {
	CreatedBy struct {
			Name string `json:"name"`
	} `json:"createdBy"`
	DateCreated	int64 `json:"dateCreated"`
	LastUpdatedBy struct {
			Name string `json:"name"`
		} `json:"lastUpdatedBy"`
	LastUpdated int64 `json:"lastUpdated"`
}