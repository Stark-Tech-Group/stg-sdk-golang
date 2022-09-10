package domain

type TagRef struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Version     int         `json:"version"`
	Ref         string      `json:"ref"`
	TargetRef   string      `json:"objectRef"`
	Enabled     bool        `json:"enabled"`
	Description interface{} `json:"description"`
	Links       struct {
		Type string `json:"type"`
		Self string `json:"self"`
	} `json:"links,omitempty"`
	Value    interface{} `json:"value"`
	Category interface{} `json:"category"`
	Tag      Tag         `json:"tag"`
	Audit    *Audit      `json:"audit,omitempty"`
}
