package entity

type QuestBase struct {
	SeedBase         `yaml:",inline"`
	Name             string `yaml:"name"`
	ScoreRankGroupID uint   `yaml:"scoreRankGroupId"`
	RewardGroupID    uint   `yaml:"rewardGroupId"`
}
