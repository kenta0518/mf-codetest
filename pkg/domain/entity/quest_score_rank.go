package entity

type QuestScoreRankType string

const (
	QuestScoreRankNone QuestScoreRankType = "none"
	QuestScoreRankC    QuestScoreRankType = "C"
	QuestScoreRankB    QuestScoreRankType = "B"
	QuestScoreRankA    QuestScoreRankType = "A"
	QuestScoreRankS    QuestScoreRankType = "S"
)

func QuestScoreRankTypeIndex(rank QuestScoreRankType) int {
	switch rank {
	case QuestScoreRankNone:
		return 0
	case QuestScoreRankC:
		return 1
	case QuestScoreRankB:
		return 2
	case QuestScoreRankA:
		return 3
	case QuestScoreRankS:
		return 4
	default:
		return -1
	}
}

type QuestScoreRank struct {
	SeedBase       `yaml:",inline"`
	GroupID        uint               `yaml:"groupId"`
	RankType       QuestScoreRankType `yaml:"rankType"`
	ScoreThreshold uint               `yaml:"scoreThreshold"`
}
