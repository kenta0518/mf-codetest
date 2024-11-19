package entity

import (
	"gorm.io/gorm"
)

type UserOshirase struct {
	gorm.Model
	UserID     uint
	OshiraseID uint
	Oshirase   Oshirase
	IsRead     bool
}

func NewUserOshirase(userID uint, Oshirase *Oshirase) *UserOshirase {
	return &UserOshirase{
		UserID:     userID,
		OshiraseID: Oshirase.ID,
		Oshirase:   *Oshirase,
		IsRead:     false,
	}
}

func (o *UserOshirase) Read() {
	o.IsRead = true
}
