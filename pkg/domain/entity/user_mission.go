package entity

import (
	"time"

	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

type UserMission struct {
	gorm.Model
	UserID     uint
	User       User
	MissionID  uint
	Mission    Mission `gorm:"foreignKey:MissionID"`
	Progress   int
	ReceivedAt *DateTime
	ResetAt    *DateTime
}

func NewUserMission(userID, missionID uint, now time.Time) *UserMission {
	return &UserMission{
		UserID:     userID,
		MissionID:  missionID,
		Progress:   0,
		ReceivedAt: nil,
		ResetAt:    &DateTime{Time: now},
	}
}

func (u UserMission) IsReceived() bool {
	return u.ReceivedAt != nil
}

func (u *UserMission) Receive(time time.Time) {
	u.ReceivedAt = &DateTime{Time: time}
}

func (u *UserMission) UpdateProgress(progress int) bool {
	if progress > u.Progress {
		u.Progress = progress
		return true
	}
	return false
}

func (u *UserMission) AddProgress(progress int) {
	if progress >= 0 {
		u.Progress += progress
	}
}

func (u UserMission) IsAchieved() bool {
	return u.Progress >= u.Mission.ConditionValue1
}

func (u *UserMission) IsInPeriod(t time.Time) bool {
	var missionTime time.Time
	if u.ResetAt != nil {
		missionTime = u.ResetAt.Time
	} else if u.CreatedAt != (time.Time{}) {
		missionTime = u.CreatedAt
	} else {
		missionTime = t
	}

	switch u.Mission.MissionGroup.GroupKind {
	case MissionGroupKindDaily:
		begin := now.With(missionTime).BeginningOfDay()
		end := now.With(missionTime).EndOfDay()
		// begin <= t < end
		return !begin.After(t) && t.Before(end)

	default:
		return true
	}
}

func (u *UserMission) Reset(now time.Time) {
	u.Progress = 0
	u.ReceivedAt = nil
	u.ResetAt = &DateTime{Time: now}
}
