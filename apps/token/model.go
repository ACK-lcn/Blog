package token

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/xid"
)

type Token struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"username" gorm:"column:username"`
	// Access token issued to the user (user needs to carry the token to access the interface)
	AccessToken string `json:"access_token"`
	// Expiration time (2h), The unit is seconds
	AccessTokenExpiredAt int `json:"access_token_expired_at"`
	// Refresh token
	RefreshToken string `json:"refresh_token"`
	// Refresh token expiration time (7d).
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at"`
	// Creation time
	CreatedAt int64 `json:"create_at"`
	// Update time
	UpdatedAt int64 `json:"update_at"`
}

func NewToken() *Token {
	return &Token{
		// The xid library can randomly generate UUID strings
		AccessToken:           xid.New().String(),
		RefreshToken:          xid.New().String(),
		CreatedAt:             time.Now().Unix(),
		AccessTokenExpiredAt:  7200,
		RefreshTokenExpiredAt: 3600 * 24 * 7,
	}
}

func (t *Token) TableName() string {
	return "tokens"
}

// Token json.Marshal serialization and format the output through the String method.
func (u *Token) String() string {
	dj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	return string(dj)
}
