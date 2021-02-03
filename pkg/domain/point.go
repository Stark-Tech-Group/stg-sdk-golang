package domain

type Points struct {
	Count	int32 		`json:"count"`
	Points 	[]*Point	`json:"points"`
}

type Point struct {
	Id 			int32 	`json:"id"`
	Ref 		string 	`json:"ref"`
	Name 		string 	`json:"site"`
	Description	string 	`json:"description"`
	Unit 		string 	`json:"unit"`
	Enabled		bool 	`json:"enabled"`
	Equip 		*Equip  `json:"equip"`
	Conn		*Conn	`json:"conn"`
	Audit		*Audit	`json:"audit"`
}
