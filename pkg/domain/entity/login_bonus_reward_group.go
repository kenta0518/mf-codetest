package entity

type LoginBonusRewardGroup struct {
	SeedBase      `yaml:",inline"`
	RewardGroupID uint `yaml:"rewardGroupId"`
	RewardContent `yaml:",inline"`
}
