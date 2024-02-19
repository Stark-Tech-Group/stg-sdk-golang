package domain

type Tag struct {
	Id    int    `json:"id,omitempty"`
	Ref   string `json:"ref,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Audit *Audit `json:"audit"`
}
