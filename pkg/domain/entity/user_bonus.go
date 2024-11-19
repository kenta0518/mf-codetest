package entity

type UserBonus struct {
	UserResourceBase
	Bonus    Bonus `gorm:"foreignKey:ResourceID"`
	Quantity int
}

func NewUserBonus(userID uint, bonusID uint, quantity int) *UserBonus {
	return &UserBonus{
		UserResourceBase: UserResourceBase{
			UserID:     userID,
			ResourceID: bonusID,
		},
		Quantity: quantity,
	}
}

func (u *UserBonus) Gain(quantity int) bool {
	if quantity < 0 {
		return false

	}
	u.Quantity += quantity
	return true
}
