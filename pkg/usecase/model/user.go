package model

import "github.com/kenta0518/mf-codetest/pkg/domain/entity"

type ProfileUser struct {
	UserID     uint   `json:"user_id"`
	ProfileUrl string `json:"profile_url"`
}

type ProfileIsNew struct {
	BannerIsNew bool `json:"banner_is_new"`
	DegreeIsNew bool `json:"degree_is_new"`
}

type TutorialResult struct {
	UserID       uint                `json:"user_id"`
	TutorialFlag entity.TutorialFlag `json:"tutorial_flag"`
}

func NewProfileUser(user *entity.User) *ProfileUser {
	return &ProfileUser{
		UserID:     user.ID,
		ProfileUrl: user.ProfileUrl,
	}
}

func NewTutorialResult(user *entity.User) *TutorialResult {
	return &TutorialResult{
		UserID:       user.ID,
		TutorialFlag: user.Property.TutorialConfirmed,
	}
}
