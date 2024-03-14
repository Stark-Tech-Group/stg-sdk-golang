package domain

type PointType struct {
	Name string `json:"name,omitempty"`
	Id   int32  `json:"id"`
}

type PointTypes struct {
	Count  int32       `json:"count"`
	Points []PointType `json:"pointTypeList"`
}
