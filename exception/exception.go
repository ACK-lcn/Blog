package exception

import (
	"fmt"
)

// Custom business exception.
type ApiException struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, format string, a ...any) *ApiException {
	return &ApiException{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
	}
}

func (e *ApiException) Error() string {
	return e.Message
}
