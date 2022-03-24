package model

type Response struct {
	Status int         `json:"status"`
	Error  bool        `json:"error"`
	Data   interface{} `json:"data"`
}
