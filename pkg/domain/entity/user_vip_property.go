package entity

import "gorm.io/gorm"

type UserVipProperty struct {
	gorm.Model
	MirrativID string `gorm:"uniqueIndex;size:256"`
	Exp        uint
	Coin       uint
}

func NewUserVipProperty(mirrativID string) *UserVipProperty {
	return &UserVipProperty{
		MirrativID: mirrativID,
	}
}

func (u *UserVipProperty) AddExp(exp uint) {
	u.Exp += exp
}

func (u *UserVipProperty) AddCoin(coin uint) {
	u.Coin += coin
}
