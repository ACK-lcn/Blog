package impl

import (
	"context"

	"github.com/ACK-lcn/Blog/apps/user"
)

// Used to explicitly constrain the implementation of the interface
var _ user.Service = &UserServiceImpl{}

type UserServiceImpl struct{}

// Create User
func (i *UserServiceImpl) CreateUser(context.Context, *user.CreateUserRequest) (*user.User, error) {
	return nil, nil
}

// Delete User
func (i *UserServiceImpl) DeleteUser(context.Context, *user.DeleteUserRequest) error {
	return nil
}

// Describe User
func (i *UserServiceImpl) DescribeUser(context.Context, *user.DescribeUserRequest) (*user.User, error) {
	return nil, nil
}
