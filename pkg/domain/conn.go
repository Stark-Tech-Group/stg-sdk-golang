package domain

type Conns struct {
	Count	int32 	`json:"count"`
	Conns 	[]*Conn	`json:"conns"`
}

type Conn struct {
	Id   	int    	`json:"id,omitempty"`
	Ref 	string 	`json:"ref,omitempty"`
	Enabled	bool 	`json:"enabled"`
	Name 	string	`json:"name,omitempty"`
	Type 	string 	`json:"type"`
	Audit	*Audit	`json:"audit"`
}