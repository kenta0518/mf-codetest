package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserRanking struct {
	gorm.Model
	UserID         uint
	RankingGroupID uint
	RankingGroup   RankingGroup
	Score          uint
	ReceivedAt     *DateTime
}

func NewUserRanking(userID, rankingGroupID uint) *UserRanking {
	return &UserRanking{
		UserID:         userID,
		RankingGroupID: rankingGroupID,
		Score:          0,
		ReceivedAt:     nil,
	}
}

func (u *UserRanking) AddScore(score uint) {
	u.Score += score
}

func (u *UserRanking) UpdateBestScore(score uint) bool {
	if u.Score >= score {
		return false
	}
	u.Score = score
	return true
}

func (u *UserRanking) IsReceived() bool {
	return u.ReceivedAt != nil
}

func (u *UserRanking) Receive(now time.Time) {
	u.ReceivedAt = &DateTime{Time: now}
}
