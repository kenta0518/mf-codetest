package entity

import (
	"time"

	"github.com/Songmu/flextime"
	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

type UserLiveState struct {
	gorm.Model
	UserID        uint
	Total         uint
	Duration      *uint
	DurationStart *time.Time
	LastLiveAt    *time.Time
}

func (u UserLiveState) HasLivedToday(user *User) bool {
	t := flextime.Now().Add(user.TimeDifference)
	if u.LastLiveAt == nil {
		return false
	}

	begin := now.With(*u.LastLiveAt).BeginningOfDay()
	end := now.With(*u.LastLiveAt).EndOfDay()
	// begin <= t < end
	return !begin.After(t) && t.Before(end)
}

func (u *UserLiveState) Live(user *User) bool {
	if u.HasLivedToday(user) {
		return false
	}

	today := flextime.Now().Add(user.TimeDifference)
	yesterday := today.AddDate(0, 0, -1)

	beginingOfDay := now.With(yesterday).BeginningOfDay()
	endOfDay := now.With(yesterday).EndOfDay()
	// beginingOfDay <= u.LastLiveAt < endOfDay
	if u.LastLiveAt != nil && !beginingOfDay.After(*u.LastLiveAt) && (*u.LastLiveAt).Before(endOfDay) {
		*u.Duration++
	} else {
		*u.Duration = 1
		u.DurationStart = &today
	}

	u.Total++
	u.LastLiveAt = &today

	return true
}
