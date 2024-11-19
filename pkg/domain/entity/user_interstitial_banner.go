package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserInterstitialBanner struct {
	gorm.Model
	UserID               uint
	User                 User
	InterstitialBannerID uint
	InterstitialBanner   InterstitialBanner
	ViewCount            uint
	LastViewedAt         *time.Time
}

func NewUserInterstitialBanner(userID uint, banner *InterstitialBanner) *UserInterstitialBanner {
	return &UserInterstitialBanner{
		UserID:               userID,
		InterstitialBannerID: banner.ID,
		InterstitialBanner:   *banner,
	}
}

func (u *UserInterstitialBanner) View(time time.Time) {
	u.ViewCount++
	u.LastViewedAt = &time
}

func (u UserInterstitialBanner) IsDisplayable(time time.Time) bool {
	if u.ViewCount >= u.InterstitialBanner.ViewCount {
		return false
	}

	// time < LastViewAt + Span
	if u.LastViewedAt != nil && time.Before(u.LastViewedAt.Add(u.InterstitialBanner.Span)) {
		return false
	}

	return true
}
