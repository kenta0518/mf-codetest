package entity

import (
	"gorm.io/gorm"
)

type RoleType string

const (
	RoleTypeMaster RoleType = "master"
	RoleTypeNone   RoleType = "none"
)

type Admin struct {
	gorm.Model
	Email    string
	Password string
	RoleType RoleType
}

func NewAdmin(email string, password string) *Admin {
	return &Admin{
		Email:    email,
		Password: password,
		RoleType: RoleTypeMaster,
	}
}
