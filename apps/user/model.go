package user

// Used to store object data stored in the database
type User struct {
	ID       int64 `json:"id"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`

	// User submitted request
	*CreateUserRequest
}
