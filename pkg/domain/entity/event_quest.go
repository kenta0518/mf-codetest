package entity

type EventQuest struct {
	QuestBase            `yaml:",inline"`
	EventID              uint `yaml:"eventId"`
	RequiredEventQuestID uint `yaml:"requiredEventQuestId"`
}
