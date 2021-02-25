package domain

type Branch struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Version     int    `json:"version"`
	Ref         string `json:"ref"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
	Links       struct {
		Type string `json:"type"`
		Self string `json:"self"`
	} `json:"links"`
	Type   string `json:"type"`
	TypeID int    `json:"typeId"`
	Parent struct {
		ID interface{} `json:"id"`
	} `json:"parent"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
	Depth     int    `json:"depth"`
	URL       string `json:"url"`
	ViewName  string `json:"viewName"`
	TargetRef string `json:"targetRef"`
	NamedPath string `json:"namedPath"`
	Hidden    bool   `json:"hidden"`
	Audit		*Audit		`json:"audit,omitempty"`
}
