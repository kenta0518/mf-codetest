package entity

type RankingKind string

const (
	RankingKindNone       RankingKind = "None"
	RankingKindHighScore  RankingKind = "HighScore"
	RankingKindTotalScore RankingKind = "TotalScore"
)

type RankingGroup struct {
	SeedBase            `yaml:",inline"`
	EventID             uint        `yaml:"eventId"`
	RankingKind         RankingKind `yaml:"rankingKind"`
	RankingRangeGroupID uint        `yaml:"rankingRangeGroupId"`
}
