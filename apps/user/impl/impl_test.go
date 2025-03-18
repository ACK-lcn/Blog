package impl_test

import (
	"context"
	"testing"

	"github.com/ACK-lcn/Blog/apps/user"
	"github.com/ACK-lcn/Blog/apps/user/impl"
)

var (
	userSvc user.Service
	ctx     = context.Background()
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "admin123"
	u, err := userSvc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: 9,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDescribeUserRequestById(t *testing.T) {
	req := user.NewDescribeUserRequestById("9")
	ins, err := userSvc.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribeUserRequestByName(t *testing.T) {
	req := user.NewDescribeUserRequestByUsername("blog")
	ins, err := userSvc.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	userSvc = &impl.UserServiceImpl{}
}
