package entity

type UserPrizeSticker struct {
	UserResourceBase
	PrizeSticker PrizeSticker `gorm:"foreignKey:ResourceID"`
	Quantity     int
}

func NewUserPrizeSticker(userID uint, prizeStickerID uint, quantity int) *UserPrizeSticker {
	return &UserPrizeSticker{
		UserResourceBase: UserResourceBase{
			UserID:     userID,
			ResourceID: prizeStickerID,
		},
		Quantity: quantity,
	}
}

func (u *UserPrizeSticker) Gain(quantity int) bool {
	if quantity < 0 {
		return false
	}
	u.Quantity += quantity
	return true
}

func (u *UserPrizeSticker) Consume(quantity int) bool {
	if u.Quantity < quantity {
		return false
	}
	u.Quantity -= quantity
	return true
}
