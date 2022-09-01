package domain

type Assets struct {
	Count  int32    `json:"count"`
	Assets []*Asset `json:"assets"`
}

type Asset struct {
	Id          int32      `json:"id,omitempty"`
	Ref         string     `json:"ref,omitempty"`
	Url        	string     `json:"url,omitempty"`
	Name        string     `json:"name,omitempty"`
	Type        string     `json:"type,omitempty"`
}