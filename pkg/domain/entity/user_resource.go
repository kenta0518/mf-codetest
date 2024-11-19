package entity

import "gorm.io/gorm"

type UserResourceBase struct {
	gorm.Model
	UserID     uint `gorm:"not null"`
	User       User `gorm:"foreignKey:UserID"`
	ResourceID uint `gorm:"not null"`
}

func (r UserResourceBase) UserResourceModule() {}

func (r UserResourceBase) GetID() uint {
	return r.ID
}

func (r UserResourceBase) IsEmpty() bool {
	return r.ID == 0
}
