package exception

func NewNotFound(format string, a ...any) *ApiException {
	return New(404, format, a...)
}
