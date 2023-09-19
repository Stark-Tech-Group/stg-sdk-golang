package domain

type CurVal struct {
	DisplayVal string `json:"displayValue"`
	Point      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Unit string `json:"unit"`
	} `json:"point"`
	Read struct {
		Urid string  `json:"urid"`
		Ts   int     `json:"ts"`
		Val  float64 `json:"val"`
	} `json:"read"`
}
