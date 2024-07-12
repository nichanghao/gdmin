package common

// PageReq 分页请求参数
type PageReq struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Limit    int
	Offset   int
}

// InitDefaultValue 初始化默认值
func (req *PageReq) InitDefaultValue() {
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	req.Limit = req.PageSize
	req.Offset = req.PageSize * (req.PageNum - 1)
}
