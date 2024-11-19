package entity

type MainQuestGroup struct {
	QuestGroupBase  `yaml:",inline"`
	RequiredQuestID uint `yaml:"requiredQuestId"`
}
