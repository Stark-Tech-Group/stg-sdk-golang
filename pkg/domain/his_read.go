package domain

type HisRead struct {
	Point struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		Unit       string `json:"unit"`
		DisplayVal string `json:"displayValue"`
	} `json:"member"`
	Start  int    `json:"start"`
	End    int    `json:"end"`
	Size   int    `json:"size"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Source string `json:"source"`
	His    []struct {
		Urid string  `json:"urid"`
		Ts   int     `json:"ts"`
		Val  float64 `json:"val"`
	} `json:"his"`
}
