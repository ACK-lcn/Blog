package exception

import (
	"fmt"
	"net/http"
)

// Custom business exception.
type ApiException struct {
	BizCode int `json:"code"`
	// Code    int    `json:"code"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
	HttpCode int    `json:"http_code"`
}

func New(code int, format string, a ...any) *ApiException {
	return &ApiException{
		BizCode:  code,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusOK,
	}
}

func (e *ApiException) Error() string {
	return e.Message
}
