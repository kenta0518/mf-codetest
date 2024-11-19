package entity

type MissionGroupKind string

const (
	MissionGroupKindPanelBeginner MissionGroupKind = "PanelBeginner"
	MissionGroupKindPanelEvent    MissionGroupKind = "PanelEvent"
	MissionGroupKindMain          MissionGroupKind = "Main"
	MissionGroupKindDaily         MissionGroupKind = "Daily"
	MissionGroupKindEvent         MissionGroupKind = "Event"
)

type MissionGroup struct {
	SeedBase  `yaml:",inline"`
	Term      `yaml:",inline"`
	Name      string           `yaml:"name"`
	GroupKind MissionGroupKind `yaml:"groupKind"`
	VipRank   uint             `yaml:"vipRank"`
}
