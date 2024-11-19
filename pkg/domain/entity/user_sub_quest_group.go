package entity

import (
	"time"

	"github.com/jinzhu/now"
)

type UserSubQuestGroup struct {
	UserResourceBase
	SubQuestGroup SubQuestGroup `gorm:"foreignkey:ResourceID"`
	UnLock        bool          `gorm:"default:false"`
	PlayCount     uint          `gorm:"not null;default:0"`
	ResetAt       *DateTime
}

func NewUserSubQuestGroup(userID, subQuestGroupID uint, unlock bool, now time.Time) *UserSubQuestGroup {
	return &UserSubQuestGroup{
		UserResourceBase: UserResourceBase{
			UserID:     userID,
			ResourceID: subQuestGroupID,
		},
		UnLock:    unlock,
		PlayCount: 0,
		ResetAt:   &DateTime{Time: now},
	}
}

func (u *UserSubQuestGroup) Unlock() {
	u.UnLock = true
}

func (u *UserSubQuestGroup) Reset(now time.Time) {
	u.PlayCount = 0
	u.ResetAt = &DateTime{Time: now}
}

func (u *UserSubQuestGroup) AddPlayCount() {
	u.PlayCount++
}

func (u *UserSubQuestGroup) IsInPeriod(t time.Time) bool {
	var groupTime time.Time
	if u.ResetAt != nil {
		groupTime = u.ResetAt.Time
	} else if u.CreatedAt != (time.Time{}) {
		groupTime = u.CreatedAt
	} else {
		return false
	}

	switch u.SubQuestGroup.PlayLimitType {
	case PlayLimitTypeDaily:
		begin := now.With(groupTime).BeginningOfDay()
		end := now.With(groupTime).EndOfDay()

		return !begin.After(t) && t.Before(end)
	case PlayLimitTypeWeekly:
		now.WeekStartDay = time.Monday
		begin := now.With(groupTime).BeginningOfWeek()
		end := now.With(groupTime).EndOfWeek()

		return !begin.After(t) && t.Before(end)

	default:
		return true
	}
}
