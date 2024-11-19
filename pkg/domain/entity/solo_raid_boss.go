package entity

type SoloRaidBoss struct {
	SeedBase            `yaml:",inline"`
	SoloRaidBossGroupID uint `yaml:"soloRaidBossGroupId"`
	Level               uint `yaml:"level"`
	HP                  uint `yaml:"hp"`
}
