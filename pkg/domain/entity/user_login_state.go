package entity

import (
	"time"

	"github.com/Songmu/flextime"
	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

type UserLoginState struct {
	gorm.Model
	UserID        uint
	Total         uint
	Duration      *uint
	DurationStart *time.Time
	LastLoginAt   *time.Time
}

func (u UserLoginState) HasLoggedInToday(user *User) bool {
	t := flextime.Now().Add(user.TimeDifference)
	if u.LastLoginAt == nil {
		return false
	}

	begin := now.With(*u.LastLoginAt).BeginningOfDay()
	end := now.With(*u.LastLoginAt).EndOfDay()
	// begin <= t < end
	return !begin.After(t) && t.Before(end)
}

func (u *UserLoginState) Login(user *User) bool {
	if u.HasLoggedInToday(user) {
		return false
	}

	today := flextime.Now().Add(user.TimeDifference)
	yesterday := today.AddDate(0, 0, -1)

	beginingOfDay := now.With(yesterday).BeginningOfDay()
	endOfDay := now.With(yesterday).EndOfDay()
	// beginingOfDay <= u.LastLoginAt < endOfDay
	if u.LastLoginAt != nil && !beginingOfDay.After(*u.LastLoginAt) && (*u.LastLoginAt).Before(endOfDay) {
		*u.Duration++
	} else {
		*u.Duration = 1
		u.DurationStart = &today
	}

	u.Total++
	u.LastLoginAt = &today

	return true
}
