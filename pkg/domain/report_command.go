package domain

// ReportCommands is a struct that represents a list of report commands
type ReportCommands struct {
	Count          int32            `json:"count"`
	ReportCommands []*ReportCommand `json:"reportCommands"`
}

// ReportCommand is a struct that represents a report command
type ReportCommand struct {
	Id        int32   `json:"id,omitempty"`
	Report    *Report `json:"report,omitempty"`
	Order     int     `json:"order,omitempty"`
	Statement string  `json:"statement,omitempty"`
}
