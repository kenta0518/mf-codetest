package entity

type CharacterEnhance struct {
	SeedBase      `yaml:",inline"`
	GroupID       uint `yaml:"groupId"`
	EnhanceLevel  uint `yaml:"enhanceLevel"`
	RewardContent `yaml:",inline"`
}
