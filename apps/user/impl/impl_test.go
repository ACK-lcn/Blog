package impl_test

import (
	"context"
	"testing"

	"github.com/ACK-lcn/Blog/apps/user"
	"github.com/ACK-lcn/Blog/apps/user/impl"
)

var (
	userSvc *impl.UserServiceImpl
	ctx     = context.Background()
)

func TestCreateUser(t *testing.T) {
	u, err := userSvc.CreateUser(ctx, &user.CreateUserRequest{})
	if err != nil {
		t.Fatal(err)
		t.Log(u)
	}
}

func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{})
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	userSvc = &impl.UserServiceImpl{}
}
