package entity

import (
	"errors"
)

var (
	ErrLackOfResources       = errors.New("lack of resources")
	ErrUserEquipmentNotFound = errors.New("user equipment not found")
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	APIKey string `json:"api_key"`
}

func (u User) IsEmpty() bool {
	return u.ID == 0
}
