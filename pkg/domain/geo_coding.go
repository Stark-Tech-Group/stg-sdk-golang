package domain

import "encoding/json"

type GeoCoding struct {
	Latitude  json.Number `json:"latitude"`
	Longitude json.Number `json:"longitude"`
	County    string      `json:"county"`
}
