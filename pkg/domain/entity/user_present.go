package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserPresent struct {
	gorm.Model
	Term
	RewardContent
	MirrativID   string `gorm:"index:idx_mirrativ_id_received_time,priority:1"`
	PresentID    *uint
	Present      *Present
	Condition    string
	ReceivedTime *time.Time `gorm:"index:idx_mirrativ_id_received_time,priority:2"`
}

func (u *UserPresent) CanBeReceived(user *User) bool {
	return u.ReceivedTime == nil && u.IsInTerm(user)
}

func (u *UserPresent) IsReceived() bool {
	return u.ReceivedTime != nil
}

func (u *UserPresent) BeReceived(now time.Time) {
	u.ReceivedTime = &now
}

func NewUserPresent(mirrativID string, reward RewardContent, cond string, term Term) *UserPresent {
	return &UserPresent{
		MirrativID:    mirrativID,
		RewardContent: reward,
		Condition:     cond,
		Term:          term,
	}
}

func NewUserPresentbyPresent(mirrativID string, present *Present) *UserPresent {
	return &UserPresent{
		MirrativID:    mirrativID,
		RewardContent: present.RewardContent,
		Present:       present,
		PresentID:     &present.ID,
		Condition:     present.Condition,
		Term:          present.Term,
	}
}
