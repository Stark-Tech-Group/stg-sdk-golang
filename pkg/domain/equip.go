package domain

type Equips struct {
	Count	int32 		`json:"count"`
	Equips 	[]*Equip	`json:"equips"`
}

type Equip struct {
	Id 			int32 		`json:"id"`
	Ref 		string 		`json:"ref"`
	Name 		string 		`json:"name"`
	Points 		[]*Point 	`json:"points"`
	Site 		*Site  		`json:"site"`
	Latitude	float64 	`json:"latitude"`
	Longitude 	float64 	`json:"longitude"`
	Enabled 	bool 		`json:"enabled"`
	Description	string 		`json:"description"`
	Conn	*Conn			`json:"conn"`
	Audit	*Audit			`json:"audit"`
}
