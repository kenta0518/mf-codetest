package entity

type CharacterItemConvert struct {
	SeedBase       `yaml:",inline"`
	ConvertGroupID uint `yaml:"convertGroupId"`
	RewardContent  `yaml:",inline"`
}
