package user

import "context"

type Service interface {
	// CreateUser creates a new user.
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// DeleteUser and all associated data.
	DeleteUser(context.Context, *DeleteUserRequest) error
}

type CreateUserRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type DeleteUserRequest struct {
	UserName string `json:"username"`
}
