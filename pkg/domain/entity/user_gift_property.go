package entity

import "gorm.io/gorm"

type UserGiftProperty struct {
	gorm.Model
	MirrativID       string `gorm:"uniqueIndex;size:256"`
	GiftEventStamina uint
}

func NewUserGiftProperty(mirrativID string) *UserGiftProperty {
	return &UserGiftProperty{
		MirrativID: mirrativID,
	}
}

func (u *UserGiftProperty) AddGiftEventStamina(stamina uint) {
	u.GiftEventStamina += stamina
}
