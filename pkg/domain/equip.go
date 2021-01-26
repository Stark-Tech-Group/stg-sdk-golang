package domain

type Equip struct {
	Name string `json:"name"`
	Ref string `json:"ref"`
	Id int32 `json:"id"`
	Points []*Point `json:"points"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Enabled bool `json:"enabled"`
	Description string `json:"description"`
}
