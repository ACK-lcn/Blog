package user

type User struct {
	ID       int   `json:"id"`
	CreateAt int64 `json:"create_at"`
	*CreateUserRequest
}
