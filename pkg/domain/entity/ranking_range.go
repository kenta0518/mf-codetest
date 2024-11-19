package entity

type RankingRange struct {
	SeedBase               `yaml:",inline"`
	RankingGroupID         uint `yaml:"groupId"`
	RankRangeStart         uint `yaml:"rankRangeStart"`
	RankRangeEnd           uint `yaml:"rankRangeEnd"`
	IsHigherRankingDisplay bool `yaml:"isHigherRankingDisplay"`
}
