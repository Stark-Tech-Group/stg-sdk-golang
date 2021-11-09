package domain

type Sites struct {
	Count	int32 	`json:"count"`
	Sites 	[]*Site `json:"sites"`
}

type Site struct{
	Name 			string 		`json:"name,omitempty"`
	Ref 			string 		`json:"ref,omitempty"`
	Id 				int32 		`json:"id,omitempty"`
	Equips 			[]*Equip 	`json:"equips"`
	Profile			*Profile 	`json:"profile"`
	Latitude  		float64 	`json:"latitude,omitempty"`
	Longitude 		float64 	`json:"longitude,omitempty"`
	Enabled 		bool 		`json:"enabled"`
	Description 	string 		`json:"description,omitempty"`
	GeoCity 		string 		`json:"geoCity,omitempty"`
	GeoStateCode 	string 		`json:"geoStateCode,omitempty"`
	GeoAddress1 	string 		`json:"geoAddress1,omitempty"`
	GeoAddress2 	string 		`json:"geoAddress2,omitempty"`
	GeoPostalCode 	string 		`json:"geoPostalCode,omitempty"`
	Audit 			*Audit		`json:"audit"`
	Conn			*Conn		`json:"conn"`
}
