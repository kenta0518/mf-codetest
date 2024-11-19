package entity

type PrizeSticker struct {
	SeedBase    `yaml:",inline"`
	Term        `yaml:",inline"`
	Name        string   `yaml:"name"`
	EffectType  ItemType `yaml:"effectType"`
	EffectValue uint     `yaml:"effectValue"`
}
