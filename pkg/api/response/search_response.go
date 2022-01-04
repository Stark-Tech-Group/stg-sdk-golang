package response

type SearchResponse struct {
	Count       int    `json:"count"`
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"pageSize"`
	QueryTime   int    `json:"queryTime"`
	Sort        string `json:"sort"`
	Order       string `json:"order"`
	Query       string `json:"query"`
	Assets      []struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		ID       int    `json:"id"`
		Enabled  bool   `json:"enabled"`
		Ref      string `json:"ref"`
		Branches []struct {
			Path string `json:"path"`
			ID   int    `json:"id"`
		} `json:"branches"`
		Asset interface{} `json:"asset"`
	} `json:"assets"`
}
