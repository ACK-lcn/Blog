package user

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	// CreateUser creates a new user.
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// DeleteUser and all associated data.
	DeleteUser(context.Context, *DeleteUserRequest) error
	// DescribeUser returns a user.
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
}

type CreateUserRequest struct {
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`

	// Suggestion: Store in the database in JSON format;
	// use the gorm library (gorm:"serializer:json") field to complete it.
	Label map[string]string `json:"label" gorm:"serializer:json"`

	// hash password
	isHashed bool
}

// Validata validates the request. Verify user parameters.
func (req *CreateUserRequest) Validata() error {
	if req.Username == "" || req.Password == "" {
		return fmt.Errorf("username or password is empty")
	}
	return nil
}

// PasswordHash encrypts the password.
func (req *CreateUserRequest) PasswordHash() {
	if req.isHashed {
		return
	}
	b, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt.GenerateFromPassword error", err)
	}
	req.Password = string(b)
	req.isHashed = true
}

func (req *CreateUserRequest) SetIsHashed() {
	req.isHashed = true
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:  ROLE_AUDITOR,
		Label: map[string]string{},
	}
}

func (req *DeleteUserRequest) IdString() string {
	return fmt.Sprintf("%d", req.Id)
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}

func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeValue: id,
	}
}

func NewDescribeUserRequestByUsername(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy:    DESCRIBE_BY_USERNAME,
		DescribeValue: username,
	}
}

// Supports ID and Username query.
type DescribeUserRequest struct {
	DescribeBy    DescribeBy `json:"describe_by"`
	DescribeValue string     `json:"describe_value"`
}
