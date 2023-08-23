package domain

import "time"

type AuditLog struct {
	DateCreated            time.Time `json:"dateCreated"`
	LastUpdated            time.Time `json:"lastUpdated"`
	TargetRef              string    `json:"targetRef"`
	Actor                  string    `json:"actor"`
	URI                    string    `json:"uri"`
	ClassName              string    `json:"className"`
	PersistedObjectId      string    `json:"persistedObjectId"`
	PersistedObjectVersion int64     `json:"persistedObjectVersion"`
	EventName              string    `json:"eventName"`
	PropertyName           string    `json:"propertyName"`
	OldValue               string    `json:"oldValue"`
	NewValue               string    `json:"newValue"`
}
