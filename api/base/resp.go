package base

type Resp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type PageData struct {
	Total int64 `json:"total"`
	Page  int   `json:"pageNum"`
	Size  int   `json:"pageSize"`
	Data  any   `json:"data"`
}
