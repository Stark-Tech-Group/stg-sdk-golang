package domain

type AssetTree struct {
	AssetTreeMeta struct {
		Size    int  `json:"size"`
		Partial bool `json:"partial"`
			} `json:"assetTreeMeta"`
	AssetTree []struct {
		TargetRef     string        `json:"targetRef"`
		Hidden        bool          `json:"hidden"`
		HasChildren   bool          `json:"hasChildren"`
		Icon          string        `json:"icon"`
		Type          string        `json:"type"`
		URL           string        `json:"url"`
		ParentID      string        `json:"parentId"`
		Path          string        `json:"path"`
		Ref           string        `json:"ref"`
		Depth         int           `json:"depth"`
		ViewName      string        `json:"viewName"`
		Children      []interface{} `json:"children"`
		Name          string        `json:"name"`
		TypeID        int           `json:"typeId"`
		ID            int           `json:"id"`
		TotalChildren int           `json:"totalChildren"`
		NamedPath     string        `json:"namedPath"`
	} `json:"assetTree"`
}
