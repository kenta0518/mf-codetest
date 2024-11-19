package entity

import "github.com/thoas/go-funk"

type ClearRewardType string

const (
	ClearRewardFirstClear ClearRewardType = "FirstClear"
	ClearRewardRankC      ClearRewardType = "ClearRankC"
	ClearRewardRankB      ClearRewardType = "ClearRankB"
	ClearRewardRankA      ClearRewardType = "ClearRankA"
	ClearRewardRankS      ClearRewardType = "ClearRankS"
)

func GetClearRewardType(oldRank, newRank QuestScoreRankType) []ClearRewardType {
	if newRank == QuestScoreRankNone || oldRank == newRank || oldRank == QuestScoreRankS {
		return []ClearRewardType{}
	}

	oldClearReward := []ClearRewardType{}
	newClearReward := []ClearRewardType{}
	switch oldRank {
	case QuestScoreRankC:
		oldClearReward = []ClearRewardType{ClearRewardRankC}
	case QuestScoreRankB:
		oldClearReward = []ClearRewardType{ClearRewardRankC, ClearRewardRankB}
	case QuestScoreRankA:
		oldClearReward = []ClearRewardType{ClearRewardRankC, ClearRewardRankB, ClearRewardRankA}
	case QuestScoreRankS:
		oldClearReward = []ClearRewardType{ClearRewardRankC, ClearRewardRankB, ClearRewardRankA, ClearRewardRankS}
	}
	switch newRank {
	case QuestScoreRankC:
		newClearReward = []ClearRewardType{ClearRewardRankC}
	case QuestScoreRankB:
		newClearReward = []ClearRewardType{ClearRewardRankC, ClearRewardRankB}
	case QuestScoreRankA:
		newClearReward = []ClearRewardType{ClearRewardRankC, ClearRewardRankB, ClearRewardRankA}
	case QuestScoreRankS:
		newClearReward = []ClearRewardType{ClearRewardRankC, ClearRewardRankB, ClearRewardRankA, ClearRewardRankS}
	}

	result := []ClearRewardType{}
	for _, new := range newClearReward {
		if !funk.Contains(oldClearReward, new) {
			result = append(result, new)
		}
	}

	return result
}

type QuestReward struct {
	SeedBase        `yaml:",inline"`
	GroupID         uint            `yaml:"groupID"`
	ClearRewardType ClearRewardType `yaml:"clearRewardType"`
	RewardContent   `yaml:",inline"`
}
