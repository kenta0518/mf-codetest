package entity

import "gorm.io/gorm"

type RectangleBanner struct {
	gorm.Model
	Term
	DisplayOrder    uint
	BannerID        string
	TransitionType  string
	TransitionValue uint
}

func (b *RectangleBanner) Update(displayOrder uint, bannerID string, transitionType string, transitionValue uint, term Term) {
	b.DisplayOrder = displayOrder
	b.BannerID = bannerID
	b.TransitionType = transitionType
	b.TransitionValue = transitionValue
	b.Term = term
}
