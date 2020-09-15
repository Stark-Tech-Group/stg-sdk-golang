package domain

type Equip struct {
	Name string     `json:"site"`
	Ref string      `json:"ref"`
	Id int32        `json:"id"`
	Points []*Point `json:"points"`
}
