package entity

type HappyBox struct {
	SeedBase      `yaml:",inline"`
	GroupId       uint `yaml:"groupId"`
	RewardContent `yaml:",inline"`
	Weight        uint `yaml:"weight"`
}
