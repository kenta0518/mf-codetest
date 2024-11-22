package entity

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID      int    `json:"user_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
