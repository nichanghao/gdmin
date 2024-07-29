package common

// PageReq 分页请求参数
type PageReq struct {
	Current int `json:"current"`
	Size    int `json:"size"`
	Limit   int
	Offset  int
}

// PageResp 分页返回参数
type PageResp struct {
	Total   int64 `json:"total"`
	Records any   `json:"records"`
	Current int   `json:"current"`
	Size    int   `json:"size"`
}

// InitDefaultValue 初始化默认值
func (req *PageReq) InitDefaultValue() {
	if req.Current == 0 {
		req.Current = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}

	req.Limit = req.Size
	req.Offset = req.Size * (req.Current - 1)
}
