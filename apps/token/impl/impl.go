package impl

import (
	"context"
	"time"

	"github.com/ACK-lcn/Blog/apps/token"
	"github.com/ACK-lcn/Blog/apps/user"
	"github.com/ACK-lcn/Blog/conf"
	"github.com/ACK-lcn/Blog/exception"
	"github.com/ACK-lcn/Blog/ioc"
	"gorm.io/gorm"
)

func Init() {
	ioc.Controller().Register(&TokenServiceImpl{})
}

type TokenServiceImpl struct {
	// db
	db *gorm.DB
	// Depends on the User function module.
	// It is strongly not recommended to directly operate the user module database (users).
	// Depends on another business area: user management area
	user user.Service
}

func (i *TokenServiceImpl) Init() error {
	i.db = conf.C().MySQL.GetConnection().Debug()
	return nil
}

func (i *TokenServiceImpl) Name() string {
	return token.AppName
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

	// Check the number of active tokens for the user.
	var activeToken []token.Token
	if err := i.db.WithContext(ctx).
		Where("user_id = ? AND Access_token_expired_at > ?", u.Id, time.Now()).
		Find(&activeToken).Error; err != nil {
		return nil, err
	}

	// If the user has more than 2 active tokens, delete  the oldest one.
	if len(activeToken) >= 2 {
		oldstToken := activeToken[0]
		for _, t := range activeToken {
			if t.CreatedAt.Before(oldstToken.CreatedAt) {
				oldstToken = t
			}
		}
		if err := i.db.WithContext(ctx).Delete(&oldstToken).Error; err != nil {
			return nil, err
		}
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
		if err == gorm.ErrRecordNotFound {
			return nil, token.TokenNotFound
		}
		return nil, err
	}

	// Check if the token is valid.
	if err := tk.IsExpired(); err != nil {
		return nil, err
	}

	return tk, nil
}

// logout interface (Destroy token)
func (i *TokenServiceImpl) Logout(ctx context.Context, req *token.LogoutRequest) error {
	tk := token.NewToken()
	err := i.db.WithContext(ctx).Where("access_token = ?", req.AccessToken).First(tk).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return token.TokenNotFound
		}
		return err
	}

	// Delete Token
	if err := i.db.WithContext(ctx).Delete(tk).Error; err != nil {
		return err
	}

	return nil
}
