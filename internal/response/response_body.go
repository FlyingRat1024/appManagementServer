package response

type ResBody struct {
	Status int         `json:"status"`
	Param  interface{} `json:"param"`
	Msg    string      `json:"msg"`
}
