package entity

type MainAdv struct {
	SeedBase            `yaml:",inline"`
	MainAdvGroupID      uint `yaml:"mainAdvGroupId"`
	MainAdvGroup        MainAdvGroup
	StoryName           string `yaml:"storyName"`
	RequiredMainQuestID uint   `yaml:"requiredMainQuestId"`
}
