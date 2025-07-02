package exception

// NotFound Status Code: 404
func NewNotFound(format string, a ...any) *ApiException {
	return New(404, format, a...)
}

func IsNotFound(err error) bool {
	if e, ok := err.(*ApiException); ok {
		if e.BizCode == 404 {
			return true
		}
	}
	return false
}

// Authentication failed Status code: 5000
func NewAuthFailed(format string, a ...any) *ApiException {
	return New(5000, format, a...)
}

func NewTokenExpired(format string, a ...any) *ApiException {
	return New(5001, format, a...)
}

func NewPermissionDeny(formmat string, a ...any) *ApiException {
	return New(5002, formmat, a...)
}
