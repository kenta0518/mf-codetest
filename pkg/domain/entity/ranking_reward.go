package entity

type RankingReward struct {
	SeedBase       `yaml:",inline"`
	RankingGroupID uint `yaml:"rankingGroupId"`
	RankingRangeID uint `yaml:"rankingRangeId"`
	RewardContent  `yaml:",inline"`
}
