package impl

import (
	"context"

	"github.com/ACK-lcn/Blog/apps/user"
	"gorm.io/gorm"
)

// Used to explicitly constrain the implementation of the interface
var _ user.Service = &UserServiceImpl{}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		db: nil,
	}
}

type UserServiceImpl struct {
	db *gorm.DB
}

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
