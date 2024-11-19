package entity

type StarRankTable struct {
	SeedBase        `yaml:",inline"`
	Rank            uint        `yaml:"rank"`
	RankMinExp      uint        `yaml:"rankMinExp"`
	RankMaxExp      uint        `yaml:"rankMaxExp"`
	ContentType     ContentType `yaml:"contentType"`
	ContentId       uint        `yaml:"contentId"`
	ContentQuantity int         `yaml:"contentQuantity"`
}

func (s *StarRankTable) GetRewardContent() RewardContent {
	return RewardContent{ContentType: s.ContentType, ContentID: s.ContentId, ContentQuantity: s.ContentQuantity}
}
