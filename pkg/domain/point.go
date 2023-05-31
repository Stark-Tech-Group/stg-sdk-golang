package domain

type Points struct {
	Count  int32    `json:"count"`
	Points []*Point `json:"points"`
}

type Point struct {
	Id          int32      `json:"id,omitempty"`
	Ref         string     `json:"ref,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	RemoteRef   string     `json:"remoteRef,omitempty"`
	Urid        *Urid      `json:"urid,omitempty"`
	PointType   *PointType `json:"pointType,omitempty"`
	Unit        string     `json:"unit"`
	Enabled     bool       `json:"enabled"`
	Equip       *Equip     `json:"equip,omitempty"`
	Conn        *Conn      `json:"conn,omitempty"`
	Audit       *Audit     `json:"audit,omitempty"`
}
