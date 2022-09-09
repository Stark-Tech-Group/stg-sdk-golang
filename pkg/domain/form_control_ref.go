package domain

type FormControlRefList struct {
	Count  int32    `json:"count"`
	FormControlRefList []*FormControlRef `json:"formControlList"`
}

type FormControlRef struct {
	Id          int32      	`json:"id,omitempty"`
	Ref         string     	`json:"ref,omitempty"`
	Name        string     	`json:"name"`
	Enabled     bool       	`json:"enabled,omitempty"`
	Description string     	`json:"description"`
	FormControl string     	`json:"formControl"`
	SortOrder 	int32     	`json:"order"`
	Key 		string     	`json:"key"`
	Value 		string     	`json:"value"`
	TargetRef 	string     	`json:"targetRef"`
	Audit       *Audit     	`json:"audit,omitempty"`
}

