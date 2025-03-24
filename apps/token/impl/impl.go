package impl

import (
	"context"

	"github.com/ACK-lcn/Blog/apps/token"
	"github.com/ACK-lcn/Blog/apps/user"
	"github.com/ACK-lcn/Blog/conf"
	"github.com/ACK-lcn/Blog/exception"
	"gorm.io/gorm"
)

type TokenServiceImpl struct {
	// db
	db *gorm.DB
	// Depends on the User function module.
	// It is strongly not recommended to directly operate the user module database (users).
	// Depends on another business area: user management area
	user user.Service
}

func NewTokenServiceImpl(userSvcImpl user.Service) *TokenServiceImpl {
	return &TokenServiceImpl{
		db:   conf.C().MySQL.GetConnection(),
		user: userSvcImpl,
		// Turn on debug
		// db: conf.C().MySQL.GetConnection().Debug(),
	}
}

// login interface (Issue token)
func (i *TokenServiceImpl) Login(ctx context.Context, req *token.LoginRequest) (*token.Token, error) {
	// Query user
	uReq := user.NewDescribeUserRequestByUsername(req.Username)
	u, err := i.user.DescribeUser(ctx, uReq)
	if err != nil {
		if exception.IsNotFound(err) {
			return nil, token.AuthFailed
		}
		return nil, err
	}

	// Compare passwords
	err = u.CheckPassword(req.Password)
	if err != nil {
		return nil, token.AuthFailed
	}

	// Issue token
	tk := token.NewToken()
	tk.UserId = u.Id
	tk.UserName = u.Username

	//  Save token
	if err := i.db.WithContext(ctx).Create(tk).Error; err != nil {
		return nil, err
	}

	return tk, nil
}

// The verification Token is used by the internal middle layer for identity verification.
// After completion, the Token is returned and the user information is obtained through the Token.
func (i *TokenServiceImpl) ValiateToken(ctx context.Context, req *token.ValiateToken) (*token.Token, error) {
	// Query token
	tk := token.NewToken()
	err := i.db.WithContext(ctx).Where("access_token = ?", req.AccessToken).First(tk).Error
	if err != nil {
		return nil, err
	}

	// Check if the token is valid.

	return nil, nil
}

// logout interface (Destroy token)
func (i *TokenServiceImpl) Logout(ctx context.Context, req *token.LogoutRequest) error {
	return nil
}
