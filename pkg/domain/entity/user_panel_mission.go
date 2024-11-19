package entity

import (
	"time"
)

type UserPanelMission struct {
	UserResourceBase
	PanelMission PanelMission `gorm:"foreignKey:ResourceID"`
	Progress     uint
	ReceivedAt   *DateTime
}

func NewUserPanelMission(userID, panelMissionID uint) *UserPanelMission {
	return &UserPanelMission{
		UserResourceBase: UserResourceBase{
			UserID:     userID,
			ResourceID: panelMissionID,
		},
	}
}

func (u *UserPanelMission) IsAchieved() bool {
	return u.Progress >= u.PanelMission.ConditionValue1
}

func (u *UserPanelMission) IsReceived() bool {
	return u.ReceivedAt != nil
}

func (u *UserPanelMission) Receive(time time.Time) {
	u.ReceivedAt = &DateTime{Time: time}
}

func (u *UserPanelMission) UpdateProgress(progress uint) bool {
	if progress > u.Progress {
		u.Progress = progress
		return true
	}
	return false
}

func (u *UserPanelMission) AddProgress(progress uint) {
	u.Progress += progress
}
