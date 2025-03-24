package token

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ACK-lcn/Blog/exception"
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
	CreatedAt time.Time `json:"create_at" gorm:"autoCreateTime"`
	// Update time
	UpdatedAt int64  `json:"update_at"`
	DeviceId  string `json:"device_id" gorm:"not null"`
}

func NewToken() *Token {
	return &Token{
		// The xid library can randomly generate UUID strings
		AccessToken:           xid.New().String(),
		RefreshToken:          xid.New().String(),
		CreatedAt:             time.Now(),
		AccessTokenExpiredAt:  7200,
		RefreshTokenExpiredAt: 3600 * 24 * 7,
	}
}

func (t *Token) TableName() string {
	return "tokens"
}

// check whether the token has expired.
func (t *Token) IsExpired() error {
	duration := time.Since(t.ExpiredTime())
	expiredSeconds := duration.Seconds()
	if expiredSeconds > 0 {
		return exception.NewTokenExpired("Token %s Expired %f Seconds", t.AccessToken, expiredSeconds)
	}
	return nil
}

// Calculate token expiration time
func (t *Token) ExpiredTime() time.Time {
	return time.Unix(t.CreatedAt.Unix(), 0).Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)
}

// Token json.Marshal serialization and format the output through the String method.
func (u *Token) String() string {
	dj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	return string(dj)
}
