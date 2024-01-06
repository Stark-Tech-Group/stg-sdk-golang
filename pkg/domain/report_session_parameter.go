package domain

// ReportSessionParameter is a struct that represents a list of report session parameters
type ReportSessionParameter struct {
	Id            int32          `json:"id,omitempty"`
	ReportSession *ReportSession `json:"reportSession,omitempty"`
	Name          string         `json:"name,omitempty"`
	Value         string         `json:"value,omitempty"`
}
