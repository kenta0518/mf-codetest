package entity

import "github.com/Songmu/flextime"

type PlayLimitType string

const (
	PlayLimitTypeNone   PlayLimitType = "None"
	PlayLimitTypeDaily  PlayLimitType = "Daily"
	PlayLimitTypeWeekly PlayLimitType = "Weekly"
	PlayLimitTypeInTerm PlayLimitType = "InTerm"
)

type SubQuestGroup struct {
	QuestGroupBase     `yaml:",inline"`
	OpenWeek           int           `yaml:"openWeek"`
	RequiredSubQuestID uint          `yaml:"requiredSubQuestId"`
	PlayLimitType      PlayLimitType `yaml:"playLimitType"`
	PlayLimit          uint          `yaml:"playLimit"`
}

func (s SubQuestGroup) IsOpen(user *User) bool {
	now := flextime.Now()
	if user != nil {
		now = flextime.Now().Add(user.TimeDifference)
	}
	weekDay := now.Weekday()
	flg := 1 << ((int(weekDay) + 6) % 7)
	return (s.OpenWeek & flg) > 0
}
