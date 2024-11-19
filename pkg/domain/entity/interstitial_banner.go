package entity

import (
	"time"

	"gorm.io/gorm"
)

type InterstitialBanner struct {
	gorm.Model
	Term
	Image           string
	Span            time.Duration
	ViewCount       uint
	TransitionType  string
	TransitionValue int
}

func (b *InterstitialBanner) Update(img string, span time.Duration, vc uint, txType string, txValue int, term Term) {
	b.Image = img
	b.Span = span
	b.ViewCount = vc
	b.TransitionType = txType
	b.TransitionValue = txValue
	b.Term = term
}
