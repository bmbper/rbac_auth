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

type PageReq[T any] struct {
	Page   int `json:"pageNum"`
	Size   int `json:"pageSize"`
	Params *T  `json:"params,omitempty"`
}

func RespOk(data any) Resp {
	return Resp{Code: "0", Msg: "success", Data: data}
}
func RespOkWithMsg(data any, msg string) Resp {
	return Resp{Code: "0", Msg: msg, Data: data}
}
func RespFail(msg string) Resp {
	return Resp{Code: "-1", Msg: msg, Data: nil}
}
func RespFailWithData(msg string, data any) Resp {
	return Resp{Code: "-1", Msg: msg, Data: data}
}
