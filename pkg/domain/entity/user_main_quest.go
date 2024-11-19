package entity

type UserMainQuest struct {
	UserResourceBase
	MainQuest MainQuest          `gorm:"foreignKey:ResourceID"`
	UnLock    bool               `gorm:"default:false"`
	IsClear   bool               `gorm:"default:false"`
	HighScore uint               `gorm:"default:0"`
	ClearRank QuestScoreRankType `gorm:"default:none"`
}

func NewUserMainQuest(userId, questId uint, unlock bool) *UserMainQuest {
	return &UserMainQuest{
		UserResourceBase: UserResourceBase{
			UserID:     userId,
			ResourceID: questId,
		},
		UnLock: unlock,
	}
}

func (q *UserMainQuest) Unlock() {
	q.UnLock = true
}

func (q *UserMainQuest) Clear() {
	q.IsClear = true
}

func (q *UserMainQuest) SetHighScore(score uint) {
	if q.HighScore < score {
		q.HighScore = score
	}
}

func (q *UserMainQuest) SetClearRank(rank QuestScoreRankType) {
	recordRank := QuestScoreRankTypeIndex(q.ClearRank)
	newRank := QuestScoreRankTypeIndex(rank)

	if recordRank < newRank {
		q.ClearRank = rank
	}
}
