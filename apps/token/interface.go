package token

import "context"

type Service interface {
	// login interface (Issue token)
	Login(context.Context, *LoginRequest) (*Token, error)
	// logout interface (Destroy token)
	Logout(context.Context, *LogoutRequest) error
}

type LoginRequest struct {
	Username string
	Password string
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

type LogoutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
