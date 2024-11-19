package entity

type CharacterCardRarity string

const (
	CharacterCardRarityN   CharacterCardRarity = "N"
	CharacterCardRarityR   CharacterCardRarity = "R"
	CharacterCardRaritySR  CharacterCardRarity = "SR"
	CharacterCardRaritySSR CharacterCardRarity = "SSR"
)

type CharacterCard struct {
	SeedBase           `yaml:",inline"`
	Term               `yaml:",inline"`
	Name               string              `yaml:"name"`
	Rarity             CharacterCardRarity `yaml:"rarity"`
	LevelMax           uint                `yaml:"levelMax"`
	ItemConvertGroupId uint                `yaml:"itemConvertGroupId"`
	EnhanceGroupID     uint                `yaml:"enhanceGroupId"`
}
