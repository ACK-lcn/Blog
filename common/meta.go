package common

import "time"

func NewMeta() *Meta {
	return &Meta{
		CreateAt: time.Now().Unix(),
		// Lable:    map[string]string{},
	}
}

// General package, storing some general information
type Meta struct {
	Id       int64 `json:"id"`
	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
	// Label    map[string]string `json:"label"`
}
