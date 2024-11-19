package entity

type VipGift struct {
	SeedBase         `yaml:",inline"`
	VipID            uint        `yaml:"vipId"`
	ContentType1     ContentType `yaml:"content1ContentType"`
	ContentId1       uint        `yaml:"content1ContentId"`
	ContentQuantity1 int         `yaml:"content1ContentQuantity"`
	ContentType2     ContentType `yaml:"content2ContentType"`
	ContentId2       uint        `yaml:"content2ContentId"`
	ContentQuantity2 int         `yaml:"content2ContentQuantity"`
	ContentType3     ContentType `yaml:"content3ContentType"`
	ContentId3       uint        `yaml:"content3ContentId"`
	ContentQuantity3 int         `yaml:"content3ContentQuantity"`
}

func (v VipGift) RewardContents() []RewardContent {
	var contents []RewardContent

	if v.ContentType1 != CONTENT_TYPE_NONE && v.ContentType1 != "" {
		content := RewardContent{v.ContentType1, v.ContentId1, v.ContentQuantity1}
		contents = append(contents, content)
	}

	if v.ContentType2 != CONTENT_TYPE_NONE && v.ContentType2 != "" {
		content := RewardContent{v.ContentType2, v.ContentId2, v.ContentQuantity2}
		contents = append(contents, content)
	}

	if v.ContentType3 != CONTENT_TYPE_NONE && v.ContentType3 != "" {
		content := RewardContent{v.ContentType3, v.ContentId3, v.ContentQuantity3}
		contents = append(contents, content)
	}

	return contents
}
