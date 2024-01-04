package response

import "net/http"

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   data,
	}
}

func NewErrorResponse(code int, message string) *Response {
	return &Response{
		Code:   code,
		Status: "error",
		Data:   message,
	}
}
