package domain

type Sites struct {
	Count	int32 	`json:"count"`
	Sites 	[]*Site `json:"sites"`
}

type Site struct{
	Name 			string 		`json:"name"`
	Ref 			string 		`json:"ref"`
	Id 				int32 		`json:"id"`
	Equips 			[]*Equip 	`json:"equips"`
	Profile			*Profile 	`json:"profile"`
	Latitude  		float64 	`json:"latitude"`
	Longitude 		float64 	`json:"longitude"`
	Enabled 		bool 		`json:"enabled"`
	Description 	string 		`json:"description"`
	GeoCity 		string 		`json:"geoCity"`
	GeoStateCode 	string 		`json:"geoStateCode"`
	GeoAddress1 	string 		`json:"geoAddress1"`
	GeoAddress2 	string 		`json:"geoAddress2"`
	GeoPostalCode 	string 		`json:"geoPostalCode"`
	Audit 			*Audit		`json:"audit"`
	Conn			*Conn		`json:"conn"`
}
