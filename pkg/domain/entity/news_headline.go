package entity

import (
	"time"

	"gorm.io/gorm"
)

type NewsHeadline struct {
	gorm.Model
	Term
	DisplayOrder uint
	Text         string
}

func NewNewsHeadline(text string, startAt, endAt, testStartAt, testEndAt time.Time, order uint) *NewsHeadline {
	return &NewsHeadline{
		Text:         text,
		Term:         *NewTerm(startAt, endAt, startAt, endAt),
		DisplayOrder: order,
	}
}

func (n *NewsHeadline) Update(text string, startAt, endAt, testStartAt, testEndAt time.Time, order uint) {
	n.Text = text
	n.Term = *NewTerm(startAt, endAt, testStartAt, testEndAt)
	n.DisplayOrder = order
}
