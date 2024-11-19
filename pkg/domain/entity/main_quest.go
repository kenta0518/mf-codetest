package entity

type MainQuest struct {
	QuestBase        `yaml:",inline"`
	MainQuestGroupID uint `yaml:"mainQuestGroupId"`
	RequiredQuestID  uint `yaml:"requiredMainQuestId"`
}
