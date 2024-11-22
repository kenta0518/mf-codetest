package model

import "github.com/kenta0518/mf-codetest/pkg/domain/entity"

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

func NewUser(user *entity.User) *User {
	return &User{
		UserID: user.ID,
		Name:   user.Name,
	}
}
