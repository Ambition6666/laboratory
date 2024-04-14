package response

type Common struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CommonData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}