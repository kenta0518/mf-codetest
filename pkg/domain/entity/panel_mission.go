package entity

type PanelMission struct {
	SeedBase            `yaml:",inline"`
	PanelMissionGroupID uint              `yaml:"panelMissionGroupId"`
	PanelMissionGroup   PanelMissionGroup `gorm:"foreignKey:PanelMissionGroupID"`
	Text                string            `yaml:"text"`
	Order               uint              `yaml:"order"`
	IsLive              bool              `yaml:"isLive"`
	Condition           MissionCondition  `yaml:"condition"`
	ConditionValue1     uint              `yaml:"conditionValue"`
	ConditionValue2     uint              `yaml:"conditionValue2"`
	RewardGroupID       uint              `yaml:"rewardGroupId"`
}
