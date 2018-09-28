package models

type ResponseInfo struct {
	ErrNO  int         `json:"err_no"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}
