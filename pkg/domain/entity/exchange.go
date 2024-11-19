package entity

type Exchange struct {
	SeedBase       `yaml:",inline"`
	Term           `yaml:",inline"`
	Name           string `yaml:"name"`
	Banner         string `yaml:"banner"`
	Order          int    `yaml:"order"`
	AppealRibon    string `yaml:"appealRibon"`
	PrizeStickerID uint   `yaml:"medalId"`
	Lineups        []ExchangeLineup
}
