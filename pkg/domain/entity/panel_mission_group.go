package entity

type PanelMissionGroup struct {
	SeedBase              `yaml:",inline"`
	Term                  `yaml:",inline"`
	UniqueID              uint             `yaml:"uniqueId"`
	Name                  string           `yaml:"name"`
	Order                 uint             `yaml:"order"`
	GroupKind             MissionGroupKind `yaml:"groupKind"`
	CompleteRewardGroupID uint             `yaml:"completeRewardGroupId"`
}
