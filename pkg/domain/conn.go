package domain

type Conns struct {
	Count	int32 	`json:"count"`
	Conns 	[]*Conn	`json:"conns"`
}

type Conn struct {
	Id   	int    	`json:"id"`
	Name 	string 	`json:"name"`
	Type 	string 	`json:"type"`
}