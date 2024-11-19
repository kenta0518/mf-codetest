package entity

import "gorm.io/gorm"

type LogBase struct {
	gorm.Model
}

func (l LogBase) LogModule() {}

func (l LogBase) GetID() uint {
	return l.ID
}

func (l LogBase) IsEmpty() bool {
	return l.ID == 0
}
