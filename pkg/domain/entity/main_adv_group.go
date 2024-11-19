package entity

type MainAdvGroup struct {
	SeedBase  `yaml:",inline"`
	Term      `yaml:",inline"`
	GroupName string `yaml:"groupName"`
}
