package entity

type SoloRaid struct {
	SeedBase              `yaml:",inline"`
	Term                  `yaml:",inline"`
	SoloRaidBossGroupID   uint `yaml:"soloRaidBossGroupId"`
	SoloRaidRewardGroupID uint `yaml:"soloRaidRewardGroupId"`
}
