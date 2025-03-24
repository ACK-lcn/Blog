package token

import "github.com/ACK-lcn/Blog/exception"

// Used to define unique exceptions for this module
var (
	AuthFailed    = exception.NewAuthFailed("用户名或者密码不正确")
	TokenNotFound = exception.NewNotFound("Token不存在")
)
