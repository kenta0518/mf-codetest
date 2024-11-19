package entity

const DEFAULT_UNLOCK_ITEM_ID = 0
const PROLOGUE_ADV_ID uint = 9999999

type UserMainAdv struct {
	UserResourceBase
	MainAdv  MainAdv `gorm:"foreignKey:ResourceID"`
	IsLock   bool
	IsReaded bool
}

func NewUserMainAdv(userID, advID uint, isLock bool) *UserMainAdv {
	return &UserMainAdv{
		UserResourceBase: UserResourceBase{
			UserID:     userID,
			ResourceID: advID,
		},
		IsLock: isLock,
	}
}

func (a *UserMainAdv) Unlock() {
	a.IsLock = false
}

func (a *UserMainAdv) Read() {
	a.IsReaded = true
}
