package entity

type ExchangeContent struct {
	SeedBase         `yaml:",inline"`
	ExchangeLineupID uint        `yaml:"groupId"`
	ContentType      ContentType `yaml:"contentType"`
	ContentID        uint        `yaml:"contentId"`
	ContentQuantity  int         `yaml:"contentQuantity"`
}

func (e ExchangeContent) RewadContent() *RewardContent {
	return &RewardContent{
		ContentType:     e.ContentType,
		ContentID:       e.ContentID,
		ContentQuantity: e.ContentQuantity,
	}
}
