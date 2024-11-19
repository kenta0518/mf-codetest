package entity

type SoloRaidReward struct {
	SeedBase              `yaml:",inline"`
	SoloRaidRewardGroupID uint `yaml:"soloRaidRewardGroupId"`
	SoloRaidBossLevel     uint `yaml:"soloRaidBossLevel"`
	RewardContent         `yaml:",inline"`
}
