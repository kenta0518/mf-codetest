package entity

import (
	"gorm.io/gorm"
)

type Present struct {
	gorm.Model
	Term
	RewardContent
	Condition         string
	IsForCurrentUsers bool
}

func NewPresent(isForCUsers bool, reward RewardContent, cond string, term Term) *Present {
	return &Present{
		Term:              term,
		RewardContent:     reward,
		Condition:         cond,
		IsForCurrentUsers: isForCUsers,
	}
}

func (r *Present) Edit(isForCUsers bool, reward RewardContent, cond string, term Term) {
	r.Term = term
	r.RewardContent = reward
	r.Condition = cond
	r.IsForCurrentUsers = isForCUsers
}

func (r *Present) CheckPresent(user *User) bool {
	if user.UserKind != SuperUser {
		return user.CreatedAt.Before(r.Term.StartAt.Time)
	} else {
		return user.CreatedAt.Before(r.Term.TestStartAt.Time)
	}
}
