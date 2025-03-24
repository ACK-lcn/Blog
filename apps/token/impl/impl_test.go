package impl_test

import (
	"context"
	"testing"

	"github.com/ACK-lcn/Blog/apps/token"
	"github.com/ACK-lcn/Blog/apps/token/impl"
	userImpl "github.com/ACK-lcn/Blog/apps/user/impl"
)

var (
	tokenSvc token.Service
	// tokenSvc token.Service
	ctx = context.Background()
)

// Test Login
func TestLogin(t *testing.T) {
	req := token.NewLoginRequest()
	req.Username = "admin"
	req.Password = "admin123"
	tk, err := tokenSvc.Login(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

// Test ValiateToken
func TestValiateToken(t *testing.T) {
	req := token.NewValiateToken("ek2m4mkmje0jbs0j3io0")
	tk, err := tokenSvc.ValiateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	// test.DevelopmentSetup()
	tokenSvc = impl.NewTokenServiceImpl(userImpl.NewUserServiceImpl())
}
