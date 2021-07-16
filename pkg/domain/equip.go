package domain

type Equips struct {
	Count	int32 		`json:"count"`
	Equips 	[]*Equip	`json:"equips"`
}

type Equip struct {
	Id 			int32 		`json:"id,omitempty"`
	Ref 		string 		`json:"ref,omitempty"`
	Name 		string 		`json:"name,omitempty"`
	Points 		[]Point 	`json:"points,omitempty"`
	Site 		Site  		`json:"site,omitempty"`
	RemoteRef	string		`json:"remoteRef,omitempty"`
	Latitude	float64 	`json:"latitude,omitempty"`
	Longitude 	float64 	`json:"longitude,omitempty"`
	Enabled 	bool 		`json:"enabled,omitempty"`
	Description	string 		`json:"description,omitempty"`
	Conn		Conn		`json:"conn,omitempty"`
	Audit		Audit		`json:"audit,omitempty"`
}
