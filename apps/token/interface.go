package token

import "context"

type Service interface {
	// login interface (Issue token)
	Login(context.Context, *LoginRequest) (*Token, error)
	// logout interface (Destroy token)
	Logout(context.Context, *LogoutRequest) error
	// The verification Token is used by the internal middle layer for identity verification.
	// After completion, the Token is returned and the user information is obtained through the Token.
	ValiateToken(context.Context, *ValiateToken) (*Token, error)
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

type ValiateToken struct {
	AccessToken string `json:"access_token"`
}
