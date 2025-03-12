package user

import "context"

type Service interface {
	// CreateUser creates a new user.
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// DeleteUser and all associated data.
	DeleteUser(context.Context, *DeleteUserRequest) error
	// DescribeUser returns a user.
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
}

type CreateUserRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	// Role     Role   `json:"role"`
}

type DeleteUserRequest struct {
	UserName string `json:"username"`
}

type DescribeUserRequest struct {
	UserName string `json:"username"`
}
