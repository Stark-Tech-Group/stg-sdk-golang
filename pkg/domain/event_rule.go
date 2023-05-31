package domain

type EventRule struct {
	Id       int    `json:"id,omitempty"`
	Ref      string `json:"ref,omitempty"`
	Enabled  bool   `json:"enabled"`
	Name     string `json:"name,omitempty"`
	RuleType string `json:"ruleType"`
	Audit    *Audit `json:"audit,omitempty"`
}
