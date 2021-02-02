package response

type SearchResponse struct {
	Count       int    `json:"count"`
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"pageSize"`
	QueryTime   int    `json:"queryTime"`
	Sort        string `json:"sort"`
	Order       string `json:"order"`
	Query       struct {
		Val     string        `json:"val"`
		Markers []struct {
			Key string      `json:"key"`
			Val interface{} `json:"val"`
		} `json:"markers"`
		Tags   	[]interface{} `json:"tags"`
		Types   []struct{
			Key string `json:"key"`
			Val string `json:"val"`
		} `json:"types"`
		Fields  []struct {
			Key string `json:"key"`
			Val string `json:"val"`
		} `json:"fields"`
	} `json:"query"`
	Assets []struct {
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
