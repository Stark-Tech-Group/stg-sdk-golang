package domain

type Profiles struct {
	Count		int32 	`json:"count"`
	Profiles 	[]*Profile `json:"profiles"`
}

type Profile struct {
	Id        	int         `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Version     int         `json:"version"`
	Ref         string      `json:"ref"`
	Enabled     bool        `json:"enabled"`
	Description	interface{} `json:"description,omitempty"`
	Links       struct {
		Type string `json:"type"`
		Self string `json:"self"`
	} `json:"links"`
	Conn	*Conn			`json:"conn"`
	Audit	*Audit			`json:"audit"`
}
