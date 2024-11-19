package entity

type UserItem struct {
	UserResourceBase
	Item     Item `gorm:"foreignKey:ResourceID"`
	Quantity int
}

func NewUserItem(userID uint, itemID uint, quantity int) *UserItem {
	return &UserItem{
		UserResourceBase: UserResourceBase{
			UserID:     userID,
			ResourceID: itemID,
		},
		Quantity: quantity,
	}
}

func (u *UserItem) Gain(quantity int) bool {
	if quantity < 0 {
		return false
	}
	u.Quantity += quantity
	return true
}

func (u *UserItem) Consume(quantity int) bool {
	if u.Quantity < quantity {
		return false
	}
	u.Quantity -= quantity
	return true
}
