package entity

type UserEventQuest struct {
	UserResourceBase
	EventQuest EventQuest         `gorm:"foreignKey:ResourceID"`
	UnLock     bool               `gorm:"default:false"`
	IsClear    bool               `gorm:"default:false"`
	HighScore  uint               `gorm:"default:0"`
	ClearRank  QuestScoreRankType `gorm:"default:none"`
}

func NewUserEventQuest(userId, questId uint, unlock bool) *UserEventQuest {
	return &UserEventQuest{
		UserResourceBase: UserResourceBase{
			UserID:     userId,
			ResourceID: questId,
		},
		UnLock:    unlock,
		IsClear:   false,
		HighScore: 0,
		ClearRank: QuestScoreRankNone,
	}
}

func (q *UserEventQuest) Unlock() {
	q.UnLock = true
}

func (q *UserEventQuest) Clear() {
	q.IsClear = true
}

func (q *UserEventQuest) SetHighScore(score uint) {
	if q.HighScore < score {
		q.HighScore = score
	}
}

func (q *UserEventQuest) SetClearRank(rank QuestScoreRankType) {
	recordRank := QuestScoreRankTypeIndex(q.ClearRank)
	newRank := QuestScoreRankTypeIndex(rank)

	if recordRank < newRank {
		q.ClearRank = rank
	}
}
