package domain

type ClimateZone struct {
	State          string `json:"state"`
	Zone           string `json:"zone"`
	Ba             string `json:"ba"`
	MoistureRegime string `json:"moistureRegime"`
}
