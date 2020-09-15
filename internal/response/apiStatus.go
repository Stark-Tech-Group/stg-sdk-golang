package response

type ApiStatus struct {
	DateTime int64  `json:"dateTime"`
	Version  string `json:"version"`
	Build    int    `json:"build"`
}


