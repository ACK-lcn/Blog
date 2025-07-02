package token

import "github.com/ACK-lcn/Blog/exception"

const (
	TOKEN_COOKIE_NAME  = "access_token"
	TOKEN_GIN_KEY_NAME = "access_token"
)

var (
	CookieNotFound = exception.NewAuthFailed("cookie %s not found", TOKEN_COOKIE_NAME)
)
