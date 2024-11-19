package entity

type VoltageRewardTable struct {
	SeedBase        `yaml:",inline"`
	GroupID         uint        `yaml:"groupId"`
	ContentType     ContentType `yaml:"contentType"`
	ContentID       uint        `yaml:"contentID"`
	ContentQuantity int         `yaml:"contentQuantity"`
	Weight          uint        `yaml:"weight"`
}
