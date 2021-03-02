package domain

type EquipCurVals struct {
	Equip struct {
		ID   int    `json:"id"`
		Ref  string `json:"ref"`
		Name string `json:"name"`
	} `json:"equip"`
	Reads []struct {
		Urid string      `json:"urid"`
		Ts   int         `json:"ts"`
		Val  float64     `json:"val"`
		Unit interface{} `json:"unit"`
	} `json:"reads"`
}