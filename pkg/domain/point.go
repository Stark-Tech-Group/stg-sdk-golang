package domain

type Points struct {
	Count	int32 		`json:"count"`
	Points 	[]*Point	`json:"points"`
}

type Point struct {
	Id 			int32 		`json:"id,omitempty"`
	Ref 		string 		`json:"ref,omitempty"`
	Name 		string 		`json:"name"`
	Description	string 		`json:"description,omitempty"`
	Urid		Urid 		`json:"urid"`
	PointType	PointType 	`json:"pointType"`
	Unit 		string 		`json:"unit"`
	Enabled		bool 		`json:"enabled"`
	Equip 		*Equip  	`json:"equip"`
	Conn		*Conn		`json:"conn"`
	Audit		*Audit		`json:"audit,omitempty"`
}
