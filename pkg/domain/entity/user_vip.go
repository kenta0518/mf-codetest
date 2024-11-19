package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserVip struct {
	gorm.Model
	UserID     uint
	VipID      uint
	Vip        Vip
	ReceivedAt *time.Time
}

func NewUserVip(userID uint, vip *Vip) *UserVip {
	return &UserVip{
		UserID: userID,
		VipID:  vip.ID,
		Vip:    *vip,
	}
}

func (u *UserVip) Receive(time time.Time) {
	u.ReceivedAt = &time
}

func (u UserVip) IsReceived() bool {
	return u.ReceivedAt != nil
}

func (u UserVip) IsAchieved(exp uint) bool {
	return u.Vip.Goal <= exp
}
