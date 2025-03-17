package user

import (
	"encoding/json"
	"fmt"

	"github.com/ACK-lcn/Blog/common"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(req *CreateUserRequest) *User {
	req.PasswordHash()

	return &User{
		Meta:              common.NewMeta(),
		CreateUserRequest: req,
	}
}

// Determine if the object is stored in the users table.
func (u *User) TableName() string {
	return "users"
}

// Use json.Marshal serialization and format the output through the String method.
func (u *User) String() string {
	dj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	return string(dj)
}

// Determine whether the user's password is correct.
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// Used to store object data stored in the database
type User struct {
	// general information
	*common.Meta
	// User submitted request
	*CreateUserRequest
}
