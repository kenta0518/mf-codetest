package entity

type Vip struct {
	SeedBase    `yaml:",inline"`
	VipRank     int    `yaml:"vip"`
	Description string `yaml:"description" gorm:"size:512"`
	Goal        uint   `yaml:"goal"`
	Effects     []VipEffect
	Gift        VipGift
}
