package entity

type SubQuest struct {
	QuestBase       `yaml:",inline"`
	SubQuestGroupID uint `yaml:"subQuestGroupId"`
	RequiredQuestID uint `yaml:"requiredQuestId"`
}
