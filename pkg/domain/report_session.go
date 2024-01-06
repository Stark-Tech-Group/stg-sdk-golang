package domain

// ReportSession is a struct that represents a report session
type ReportSession struct {
	Id           int32   `json:"id,omitempty"`
	Report       *Report `json:"report,omitempty"`
	CreatedById  int64   `json:"createdById,omitempty"`
	DateCreated  int64   `json:"dateCreated,omitempty"`
	Guid         string  `json:"guid,omitempty"`
	NumberOfRows int64   `json:"numberOfRows,omitempty"`
	TimeTakenMs  int64   `json:"timeTakenMs,omitempty"`
}
