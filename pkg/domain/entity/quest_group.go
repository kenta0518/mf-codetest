package entity

type QuestGroupBase struct {
	SeedBase `yaml:",inline"`
	Term     `yaml:",inline"`
	Name     string `yaml:"name"`
}
