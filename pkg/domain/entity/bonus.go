package entity

type BonusEffectType string

type Bonus struct {
	SeedBase        `yaml:",inline"`
	Name            string          `yaml:"name"`
	BonusEffectType BonusEffectType `yaml:"bonusEffectType"`
	EffectValue     uint            `yaml:"effectValue"`
}
