package impl_test

import (
	"context"

	"github.com/ACK-lcn/Blog/apps/token/impl"
	userImpl "github.com/ACK-lcn/Blog/apps/user/impl"
)

var (
	tokenSvc *impl.TokenServiceImpl
	// tokenSvc token.Service
	ctx = context.Background()
)

func init() {
	// test.DevelopmentSetup()
	tokenSvc = impl.NewTokenServiceImpl(userImpl.NewUserServiceImpl())
}
