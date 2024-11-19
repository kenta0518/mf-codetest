package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserExchangeLineup struct {
	gorm.Model
	UserID           uint
	ExchangeLineupID uint
	ExchangeLineup   ExchangeLineup
	Stock            int
	ResetAt          *time.Time
}

func NewUserExchangeLineup(userId uint, exchangeLineup *ExchangeLineup) *UserExchangeLineup {
	return &UserExchangeLineup{
		UserID:           userId,
		ExchangeLineupID: exchangeLineup.ID,
		ExchangeLineup:   *exchangeLineup,
		Stock:            exchangeLineup.SalesLimit,
	}
}

// 消費
func (u *UserExchangeLineup) ConsumeStock(quantity int) bool {
	// 0の場合は無制限
	if u.ExchangeLineup.SalesLimit == 0 {
		return true
	}

	if u.Stock < quantity {
		return false
	}
	u.Stock -= quantity
	return true
}

func (u *UserExchangeLineup) Reset(t time.Time) bool {
	u.Stock = u.ExchangeLineup.SalesLimit
	u.ResetAt = &t

	return true
}
