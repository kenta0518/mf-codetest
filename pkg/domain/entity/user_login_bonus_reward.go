package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserLoginBonusReward struct {
	gorm.Model
	UserID             uint
	LoginBonusRewardID uint
	LoginBonusReward   LoginBonusReward
	ReceivedAt         *time.Time
	ResetAt            *time.Time
}

func NewUserLoginBonusReward(userID, rewardID uint) *UserLoginBonusReward {
	return &UserLoginBonusReward{
		UserID:             userID,
		LoginBonusRewardID: rewardID,
	}
}

func (u UserLoginBonusReward) IsReceived() bool {
	return u.ReceivedAt != nil
}

func (u *UserLoginBonusReward) Receive(time time.Time) {
	u.ReceivedAt = &time
}

func (u *UserLoginBonusReward) Reset(time time.Time) {
	u.ReceivedAt = nil
	u.ResetAt = &time
}
