package domain

type Urids struct {
	Count int32   `json:"count"`
	Urids []*Urid `json:"pointUrids"`
}

type Urid struct {
	Name string `json:"name,omitempty"`
	Id   int32  `json:"id,omitempty"`
}
