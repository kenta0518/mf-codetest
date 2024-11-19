package entity

type UserMainQuestGroup struct {
	UserResourceBase
	MainQuestGroup MainQuestGroup `gorm:"foreignKey:ResourceID"`
	UnLock         bool           `gorm:"default:false"`
}

func NewUserQuestGroup(userId, questGroupId uint, unlock bool) *UserMainQuestGroup {
	return &UserMainQuestGroup{
		UserResourceBase: UserResourceBase{
			UserID:     userId,
			ResourceID: questGroupId,
		},
		UnLock: unlock,
	}
}

func (q *UserMainQuestGroup) Unlock() {
	q.UnLock = true
}
