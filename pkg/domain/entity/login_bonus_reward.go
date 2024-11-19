package entity

type LoginBonusReward struct {
	SeedBase      `yaml:",inline"`
	LoginBonusID  uint       `yaml:"loginBonusId"`
	LoginBonus    LoginBonus `gorm:"foreignKey:LoginBonusID"`
	DayNumber     int        `yaml:"dayNumber"`
	RewardGroupID uint       `yaml:"rewardGroupId"`
}
