package domain

type MsgResp struct {
	Path  string  `json:"path"`
	Value float64 `json:"value"`
	Units string  `json:"units"`
}
