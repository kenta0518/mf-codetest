package entity

import (
	"time"

	"github.com/Songmu/flextime"
)

type ExchangeLog struct {
	LogBase
	UserID           uint `gorm:"index:idx_user_id_exchange_id,priority:1"`
	User             User
	ExchangeID       uint `gorm:"index:idx_user_id_exchange_id,priority:2"`
	Exchange         Exchange
	ExchangeLineupID uint
	ExchangeLineup   ExchangeLineup
	Quantity         int
	ExchangedAt      time.Time
}

func NewExchangeLog(UserID uint, exchangeID uint, exchangeLineupID uint, quantity int) *ExchangeLog {
	return &ExchangeLog{
		UserID:           UserID,
		ExchangeID:       exchangeID,
		ExchangeLineupID: exchangeLineupID,
		Quantity:         quantity,
		ExchangedAt:      flextime.Now(),
	}
}
