package token

type Token struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	// Access token issued to the user (user needs to carry the token to access the interface)
	AccessToken string `json:"access_token"`
	// Expiration time (2h), The unit is seconds
	AccessTokenExpiredAt int `json:"access_token_expired_at"`
	// Refresh token
	RefreshToken string `json:"refresh_token"`
	// Refresh token expiration time (7d).
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at"`
	// Creation time
	CreatedAt int `json:"create_at"`
	// Update time
	UpdatedAt int `json:"update_at"`
}
