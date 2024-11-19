package entity

type GachaPrize struct {
	SeedBase      `yaml:",inline"`
	GachaID       uint `yaml:"gachaId"`
	RewardContent `yaml:",inline"`
	Weight        uint `yaml:"weight"`
	SpecialWeight uint `yaml:"specialWeight"`
	IsPickUp      bool `yaml:"isPickUp"`
}
