package domain

type EquipTypes struct {
	Count	int32 		`json:"count"`
	EquipTypes 	[]EquipType	`json:"equipTypeList"`
}

type EquipType struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}
