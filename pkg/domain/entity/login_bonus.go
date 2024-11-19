package entity

type LoginBonusType string

const (
	LoginBonusTypeDaily LoginBonusType = "Daily"
	LoginBonusTypeEvent LoginBonusType = "Event"
)

type LoginBonus struct {
	SeedBase       `yaml:",inline"`
	Term           `yaml:",inline"`
	LoginBonusType LoginBonusType `yaml:"loginBonusKind"`
	VipRank        uint           `yaml:"vipRank"`
	Order          uint           `yaml:"order"`
}
