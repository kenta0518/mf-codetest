package entity

type UserSubQuest struct {
	UserResourceBase
	SubQuest  SubQuest           `gorm:"foreignKey:ResourceID"`
	UnLock    bool               `gorm:"default:false"`
	IsClear   bool               `gorm:"default:false"`
	HighScore uint               `gorm:"default:0"`
	ClearRank QuestScoreRankType `gorm:"default:none"`
}

func NewUserSubQuest(userId, subQuestId uint, unlock bool) *UserSubQuest {
	return &UserSubQuest{
		UserResourceBase: UserResourceBase{
			UserID:     userId,
			ResourceID: subQuestId,
		},
		UnLock:    unlock,
		IsClear:   false,
		HighScore: 0,
		ClearRank: QuestScoreRankNone,
	}
}

func (u *UserSubQuest) Unlock() {
	u.UnLock = true
}

func (u *UserSubQuest) Clear() {
	u.IsClear = true
}

func (u *UserSubQuest) SetHighScore(score uint) {
	if u.HighScore < score {
		u.HighScore = score
	}
}

func (u *UserSubQuest) SetClearRank(rank QuestScoreRankType) {
	recordRank := QuestScoreRankTypeIndex(u.ClearRank)
	newRank := QuestScoreRankTypeIndex(rank)

	if recordRank < newRank {
		u.ClearRank = rank
	}
}
