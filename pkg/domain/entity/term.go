package entity

import (
	"time"

	"github.com/Songmu/flextime"
)

type Term struct {
	StartAt     DateTime `yaml:"startAt"`
	EndAt       DateTime `yaml:"endAt"`
	TestStartAt DateTime `yaml:"testStartAt"`
	TestEndAt   DateTime `yaml:"testEndAt"`
}

func NewTerm(start, end, testStartAt, testEndAt time.Time) *Term {
	return &Term{
		StartAt:     DateTime{start},
		EndAt:       DateTime{end},
		TestStartAt: DateTime{testStartAt},
		TestEndAt:   DateTime{testEndAt},
	}
}

func (t *Term) IsInTerm(user *User) bool {
	if user == nil {
		now := flextime.Now()
		return t.StartAt.Before(now) && t.EndAt.After(now)
	}
	now := flextime.Now().Add(user.TimeDifference)
	if user.IsSuperUser() {
		return t.TestStartAt.Before(now) && t.TestEndAt.After(now)
	}
	return t.StartAt.Before(now) && t.EndAt.After(now)
}

func (t *Term) IsEndTerm(user *User) bool {
	if user == nil {
		now := flextime.Now()
		return t.EndAt.Before(now)
	}
	now := flextime.Now().Add(user.TimeDifference)
	if user.IsSuperUser() {
		return t.TestEndAt.Before(now)
	}

	return t.EndAt.Before(now)
}
