package entity

type MissionReward struct {
	SeedBase      `yaml:",inline"`
	GroupID       uint `yaml:"groupId"`
	RewardContent `yaml:",inline"`
}
