package common

// General package, storing some general information
type Meta struct {
	ID       int64 `json:"id"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}
