package domain

type Profiles struct {
	Count int32
	Profiles []*Profile `json:"profiles"`
}

type Profile struct {
	Id        	int         `json:"id"`
	Name        string      `json:"name"`
	Version     int         `json:"version"`
	Ref         string      `json:"ref"`
	Enabled     bool        `json:"enabled"`
	Description	interface{} `json:"description"`
	Links       struct {
		Type string `json:"type"`
		Self string `json:"self"`
	} `json:"links"`
	RemoteRef 	interface{} `json:"remoteRef"`
	Conn      	struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"conn"`
	Audit struct {
		CreatedBy struct {
			Name string `json:"name"`
		} `json:"createdBy"`
		DateCreated   int64 `json:"dateCreated"`
		LastUpdatedBy struct {
			Name string `json:"name"`
		} `json:"lastUpdatedBy"`
		LastUpdated int64 `json:"lastUpdated"`
	} `json:"audit"`
}
