package domain

type Equips struct {
	Count	int32 		`json:"count"`
	Equips 	[]*Equip	`json:"equips"`
}

type Equip struct {
	Id 			int32 		`json:"id,omitempty"`
	Ref 		string 		`json:"ref,omitempty"`
	Name 		string 		`json:"name"`
	Points 		[]*Point 	`json:"points"`
	Site 		*Site  		`json:"site"`
	RemoteRef	string		`json:"remoteRef,omitempty"`
	Latitude	float64 	`json:"latitude,omitempty"`
	Longitude 	float64 	`json:"longitude,omitempty"`
	Enabled 	bool 		`json:"enabled"`
	Description	string 		`json:"description,omitempty"`
	Conn	*Conn			`json:"conn"`
	Audit	*Audit			`json:"audit"`
}
