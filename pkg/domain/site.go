package domain


type Site struct{

	Name string     `json:"site"`
	Ref string      `json:"ref"`
	Id int32        `json:"id"`
	Equips []*Equip `json:"equips"`
}
