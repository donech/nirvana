package v1

type Pager struct {
	Cursor  int64 `json:"cursor"`
	Size    int64 `json:"size"`
	HasMore bool  `json:"has_more"`
}
